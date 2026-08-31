# KeyDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Hsms** | **[]string** |  | 
**Keys** | [**[]Key**](Key.md) |  | 

## Methods

### NewKeyDomain

`func NewKeyDomain(id string, hsms []string, keys []Key, ) *KeyDomain`

NewKeyDomain instantiates a new KeyDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDomainWithDefaults

`func NewKeyDomainWithDefaults() *KeyDomain`

NewKeyDomainWithDefaults instantiates a new KeyDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyDomain) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *KeyDomain) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *KeyDomain) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *KeyDomain) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *KeyDomain) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHsms

`func (o *KeyDomain) GetHsms() []string`

GetHsms returns the Hsms field if non-nil, zero value otherwise.

### GetHsmsOk

`func (o *KeyDomain) GetHsmsOk() (*[]string, bool)`

GetHsmsOk returns a tuple with the Hsms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsms

`func (o *KeyDomain) SetHsms(v []string)`

SetHsms sets Hsms field to given value.


### GetKeys

`func (o *KeyDomain) GetKeys() []Key`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *KeyDomain) GetKeysOk() (*[]Key, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *KeyDomain) SetKeys(v []Key)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


