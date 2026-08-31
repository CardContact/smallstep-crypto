# SignatureInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Algo** | [**KeyAlgorithm**](KeyAlgorithm.md) |  | 

## Methods

### NewSignatureInput

`func NewSignatureInput(hash string, algo KeyAlgorithm, ) *SignatureInput`

NewSignatureInput instantiates a new SignatureInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureInputWithDefaults

`func NewSignatureInputWithDefaults() *SignatureInput`

NewSignatureInputWithDefaults instantiates a new SignatureInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SignatureInput) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SignatureInput) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SignatureInput) SetHash(v string)`

SetHash sets Hash field to given value.


### GetAlgo

`func (o *SignatureInput) GetAlgo() KeyAlgorithm`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *SignatureInput) GetAlgoOk() (*KeyAlgorithm, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *SignatureInput) SetAlgo(v KeyAlgorithm)`

SetAlgo sets Algo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


