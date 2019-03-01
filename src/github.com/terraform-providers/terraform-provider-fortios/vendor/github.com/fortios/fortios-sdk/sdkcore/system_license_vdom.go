package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	// "strconv"
)

// JSONSystemLicenseVDOM contains the parameters for Create and Update API function
type JSONSystemLicenseVDOM struct {
	License       string `json:"license"`
}

// JSONCreateSystemLicenseVDOMOutput contains the output results for Create API function
type JSONCreateSystemLicenseVDOMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemLicenseVDOMOutput contains the output results for Update API function
// Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemLicenseVDOMOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemLicenseVDOM API operation for FortiOS creates a new ----------------
// Returns the index value of the ---------------- and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLicenseVDOM(params *JSONSystemLicenseVDOM) (output *JSONCreateSystemLicenseVDOMOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/registration/vdom/add-license"
	output = &JSONCreateSystemLicenseVDOMOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				err = fmt.Errorf("cannot get the right response")
				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateSystemLicenseVDOM API operation for FortiOS updates the specified ----------------
// Returns the index value of the ---------------- and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLicenseVDOM(params *JSONSystemLicenseVDOM, mkey string) (output *JSONUpdateSystemLicenseVDOMOutput, err error) {
	// HTTPMethod := "PUT"
	// path := "/api/v2/monitor/registration/vdom/add-license"
	// path += "/" + mkey
	// output = &JSONUpdateSystemLicenseVDOMOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// req := c.NewRequest(HTTPMethod, path, nil, bytes)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["vdom"] != nil {
	// 		output.Vdom = result["vdom"].(string)
	// 	}
	// 	if result["mkey"] != nil {
	// 		output.Mkey = result["mkey"].(string)
	// 	}
	// 	if result["status"] != nil {
	// 		if result["status"] != "success" {
	// 			err = fmt.Errorf("cannot get the right response")
	// 			return
	// 		}
	// 		output.Status = result["status"].(string)
	// 	} else {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// 	if result["http_status"] != nil {
	// 		output.HTTPStatus = result["http_status"].(float64)
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// DeleteSystemLicenseVDOM API operation for FortiOS deletes the specified ----------------
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLicenseVDOM(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/monitor/registration/vdom/add-license"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// 	if result["status"] != "success" {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemLicenseVDOM API operation for FortiOS gets the ----------------
// with the specified index value.
// Returns the requested ---------------- when the request executes successfully.
// Returns error for service API and SDK errors.
// See the ---------------- chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLicenseVDOM(mkey string) (output *JSONSystemLicenseVDOM, err error) {
	// HTTPMethod := "GET"
	// path := "/api/v2/cmdb/registration/vdom/add-license"
	// path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios reading response: %s", string(body))

	// output = &JSONSystemLicenseVDOM{}
	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
		
	// 	if result["status"] != "success" {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}

	// 	mapTmp := (result["results"].([]interface {}))[0].(map[string]interface {})

	// 	if mapTmp == nil {
	// 		return
	// 	}

	// 	if mapTmp["license"] != nil {
	// 		output.License = mapTmp["license"].(string)
	// 	}

	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}
