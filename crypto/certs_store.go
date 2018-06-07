/*
Copyright ArxanFintech Technology Ltd. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crypto

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/arxanchain/sdk-go-common/crypto/ecc"
	"github.com/arxanchain/sdk-go-common/crypto/ecc/primitives"
	"github.com/arxanchain/sdk-go-common/crypto/rsa"
	"github.com/arxanchain/sdk-go-common/log"
	"github.com/arxanchain/sdk-go-common/structs/pki"
)

var (
	logger                     = log.MustGetLogger("crypto")
	rwmutex      *sync.RWMutex = new(sync.RWMutex)
	g_CertsStore *CertsStore
)

type ICryptoLib interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) error
	Encrypt(data []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

type CertGroup struct {
	EnrollmentID   string
	PrivateKeyFile string
	PeerCertFile   string
	cryptoLib      pki.ICryptoLib
}

type CertsStore struct {
	path     string
	certsMap map[string]*CertGroup
}

// NewCertsStore New and initialize the certs store, must be called before using crypto library
// parameters:
//   path: path to the certs store witch contains all the enrollment id server-side certificates
//
// Certs Store Dir Structure:
// - root-dir:
//		 tls:
//		     tls.key
//	     enrollmentID1:
//		     enrollmentID1.cert
//	     enrollmentID2:
//		     enrollmentID2.cert
//
func NewCertsStore(path string) (certsStore *CertsStore, err error) {
	rwmutex.Lock()
	defer rwmutex.Unlock()

	// if path is new, we need create a new CertStore
	if (g_CertsStore == nil) || (g_CertsStore != nil && g_CertsStore.path != path) {
		g_CertsStore = &CertsStore{}
		g_CertsStore.path = path
		g_CertsStore.certsMap = make(map[string]*CertGroup)

		err = primitives.InitSecurityLevel(g_HashAlgorithm, g_SecurityLevel)
		if err != nil {
			logger.Errorf("Failed to set security level: [%s]", err)
			return
		}

		err = g_CertsStore.Reload()
		if err != nil {
			return
		}
	}

	return g_CertsStore, nil
}

func (this *CertsStore) Reload() error {
	// Check the root dir
	if _, err := os.Stat(this.path); os.IsNotExist(err) {
		logger.Errorf("cert store path not existing: %v", err)
		return err
	}

	// Check the tls cert dir
	tlsPath := filepath.Join(this.path, "tls")
	if _, err := os.Stat(tlsPath); os.IsNotExist(err) {
		logger.Errorf("tls cert path not existing: %v", err)
		return err
	}

	// Check the users cert dir
	usersPath := filepath.Join(this.path, "users")
	if _, err := os.Stat(usersPath); os.IsNotExist(err) {
		logger.Errorf("tls cert path not existing: %v", err)
		return err
	}

	files, err := ioutil.ReadDir(usersPath)
	if err != nil {
		logger.Errorf("open users certs store path error: %v", err)
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if _, ok := this.certsMap[f.Name()]; !ok {
				enrollmentID := f.Name()

				var privateKeyFile, peerCertFile string

				switch g_ServerClientMode {
				case SERVER_MODE:
					privateKeyFile = filepath.Join(tlsPath, "tls.key")
					peerCertFile = filepath.Join(usersPath, enrollmentID, fmt.Sprintf("%s.cert", enrollmentID))
				case CLIENT_MODE:
					privateKeyFile = filepath.Join(usersPath, enrollmentID, fmt.Sprintf("%s.key", enrollmentID))
					peerCertFile = filepath.Join(tlsPath, "tls.cert")
				default:
					return fmt.Errorf("invalide mode: %d", g_ServerClientMode)
				}

				if _, err = os.Stat(privateKeyFile); err != nil {
					logger.Warning("Stat %s error: %v", privateKeyFile, err)
					continue
				}
				if _, err = os.Stat(peerCertFile); err != nil {
					logger.Warning("Stat %s error: %v", peerCertFile, err)
					continue
				}

				certGroup := new(CertGroup)
				certGroup.EnrollmentID = enrollmentID
				certGroup.PrivateKeyFile = privateKeyFile
				certGroup.PeerCertFile = peerCertFile
				switch g_EncryptType {
				case RSA_TYPE:
					certGroup.cryptoLib, err = rsa.NewRSACryptoLib(privateKeyFile, peerCertFile)
				case ECC_TYPE:
					certGroup.cryptoLib, err = ecc.NewECCCryptoLib(privateKeyFile, peerCertFile)
				default:
					return fmt.Errorf("invalide encrypt mode: %d", g_EncryptType)
				}
				if err != nil {
					logger.Errorf("Failed to create crypto lib[%d - %d]: %v", g_ServerClientMode, g_EncryptType, err)
					os.RemoveAll(filepath.Join(usersPath, enrollmentID))
					continue
				}

				this.certsMap[enrollmentID] = certGroup

				logger.Debugf("Load keys and certs for %s done.", enrollmentID)
			}
		}
	}
	return nil
}

// GetCryptoLib can get the specific enrollment id's CryptoLib instance,
// then you can invoke following API use this Cryptolib instance:
//
//	Sign(data []byte) ([]byte, error)
//	Verify(data []byte, sig []byte) error
//	Encrypt(data []byte) ([]byte, error)
//	Decrypt(ciphertext []byte) ([]byte, error)
//
func GetCryptoLib(name string) (cryptoLib pki.ICryptoLib, err error) {
	rwmutex.Lock()
	defer rwmutex.Unlock()

	if g_CertsStore == nil {
		return nil, fmt.Errorf("certs store not be initialized")
	}

	logger.Debugf("Get cryptolib for %s", name)

	certGroup, ok := g_CertsStore.certsMap[name]
	if !ok || certGroup == nil {
		return nil, fmt.Errorf("not found certs info for %s", name)
	}

	if certGroup.cryptoLib == nil {
		switch g_EncryptType {
		case RSA_TYPE:
			certGroup.cryptoLib, err = rsa.NewRSACryptoLib(certGroup.PrivateKeyFile, certGroup.PeerCertFile)
		case ECC_TYPE:
			certGroup.cryptoLib, err = ecc.NewECCCryptoLib(certGroup.PrivateKeyFile, certGroup.PeerCertFile)
		default:
			return nil, fmt.Errorf("invalide encrypt mode: %d", g_EncryptType)
		}
		if err != nil {
			logger.Errorf("Failed to create crypto lib[%d - %d]: %v", g_ServerClientMode, g_EncryptType, err)
			return nil, err
		}
	}

	return certGroup.cryptoLib, nil
}
