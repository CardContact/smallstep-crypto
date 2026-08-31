# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**Type** | [**KeyType**](KeyType.md) |  | 
**Size** | **int32** |  | 
**KeyUseCounter** | Pointer to **int32** |  | [optional] 
**Algorithms** | Pointer to [**[]KeyAlgorithm**](KeyAlgorithm.md) |  | [optional] 
**KeyDomain** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Pubkey** | Pointer to **string** |  | [optional] 
**Hsms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKey

`func NewKey(id string, label string, type_ KeyType, size int32, ) *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Key) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Key) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Key) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *Key) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Key) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Key) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *Key) GetType() KeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Key) GetTypeOk() (*KeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Key) SetType(v KeyType)`

SetType sets Type field to given value.


### GetSize

`func (o *Key) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Key) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Key) SetSize(v int32)`

SetSize sets Size field to given value.


### GetKeyUseCounter

`func (o *Key) GetKeyUseCounter() int32`

GetKeyUseCounter returns the KeyUseCounter field if non-nil, zero value otherwise.

### GetKeyUseCounterOk

`func (o *Key) GetKeyUseCounterOk() (*int32, bool)`

GetKeyUseCounterOk returns a tuple with the KeyUseCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUseCounter

`func (o *Key) SetKeyUseCounter(v int32)`

SetKeyUseCounter sets KeyUseCounter field to given value.

### HasKeyUseCounter

`func (o *Key) HasKeyUseCounter() bool`

HasKeyUseCounter returns a boolean if a field has been set.

### GetAlgorithms

`func (o *Key) GetAlgorithms() []KeyAlgorithm`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *Key) GetAlgorithmsOk() (*[]KeyAlgorithm, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *Key) SetAlgorithms(v []KeyAlgorithm)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *Key) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetKeyDomain

`func (o *Key) GetKeyDomain() string`

GetKeyDomain returns the KeyDomain field if non-nil, zero value otherwise.

### GetKeyDomainOk

`func (o *Key) GetKeyDomainOk() (*string, bool)`

GetKeyDomainOk returns a tuple with the KeyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDomain

`func (o *Key) SetKeyDomain(v string)`

SetKeyDomain sets KeyDomain field to given value.

### HasKeyDomain

`func (o *Key) HasKeyDomain() bool`

HasKeyDomain returns a boolean if a field has been set.

### GetCert

`func (o *Key) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *Key) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *Key) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *Key) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetPubkey

`func (o *Key) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Key) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Key) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *Key) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetHsms

`func (o *Key) GetHsms() []string`

GetHsms returns the Hsms field if non-nil, zero value otherwise.

### GetHsmsOk

`func (o *Key) GetHsmsOk() (*[]string, bool)`

GetHsmsOk returns a tuple with the Hsms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsms

`func (o *Key) SetHsms(v []string)`

SetHsms sets Hsms field to given value.

### HasHsms

`func (o *Key) HasHsms() bool`

HasHsms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


