/**
 * Licensed Materials - Property of Arxan Fintech
 *
 * (C) Copyright Arxan Fintech. 2018 All Rights Reserved
 *
 * Contributors:
 *    HHH - Initial implementation
**/

package wallet

import time "time"

func (m *QueryUTXOResponse) Len() int {
	return len(m.Utxos)
}

func (m *QueryUTXOResponse) Swap(i, j int) {
	m.Utxos[i], m.Utxos[j] = m.Utxos[j], m.Utxos[i]
}

func (m *QueryUTXOResponse) Less(i, j int) bool {
	iTime := time.Unix(m.Utxos[i].CreatedAt.Time.GetSeconds(), int64(m.Utxos[i].CreatedAt.Time.GetNanos()))
	jTime := time.Unix(m.Utxos[j].CreatedAt.Time.GetSeconds(), int64(m.Utxos[j].CreatedAt.Time.GetNanos()))
	return iTime.UnixNano() < jTime.UnixNano()
}
