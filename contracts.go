package gotezos

import "encoding/json"

// GetContractStorage gets the contract storage for a contract
func (gt *GoTezos) GetContractStorage(contract string) ([]byte, error) {
	get := "/chains/main/blocks/head/context/contracts/" + contract + "/storage"
	resp, err := gt.GetResponse(get, "{}")
	if err != nil {
		return resp.Bytes, err
	}
	return resp.Bytes, nil
}

// GetContractBigMap gets a contract big map value for a contract
func (gt *GoTezos) GetContractBigMap(contract string, args MichelineExpr) ([]byte, error) {
	post := "/chains/main/blocks/head/context/contracts/" + contract + "/big_map_get"
	binArgs, err := json.Marshal(args)
	resp, err := gt.PostResponse(post, string(binArgs))
	if err != nil {
		return resp.Bytes, err
	}
	return resp.Bytes, nil
}
