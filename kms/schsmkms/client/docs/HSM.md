# HSM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DefaultKeyDomain** | **string** |  | 
**KeyDomains** | [**[]KeyDomainStatus**](KeyDomainStatus.md) |  | 
**Keys** | [**[]Key**](Key.md) |  | 

## Methods

### NewHSM

`func NewHSM(id string, defaultKeyDomain string, keyDomains []KeyDomainStatus, keys []Key, ) *HSM`

NewHSM instantiates a new HSM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMWithDefaults

`func NewHSMWithDefaults() *HSM`

NewHSMWithDefaults instantiates a new HSM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HSM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HSM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HSM) SetId(v string)`

SetId sets Id field to given value.


### GetDefaultKeyDomain

`func (o *HSM) GetDefaultKeyDomain() string`

GetDefaultKeyDomain returns the DefaultKeyDomain field if non-nil, zero value otherwise.

### GetDefaultKeyDomainOk

`func (o *HSM) GetDefaultKeyDomainOk() (*string, bool)`

GetDefaultKeyDomainOk returns a tuple with the DefaultKeyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyDomain

`func (o *HSM) SetDefaultKeyDomain(v string)`

SetDefaultKeyDomain sets DefaultKeyDomain field to given value.


### GetKeyDomains

`func (o *HSM) GetKeyDomains() []KeyDomainStatus`

GetKeyDomains returns the KeyDomains field if non-nil, zero value otherwise.

### GetKeyDomainsOk

`func (o *HSM) GetKeyDomainsOk() (*[]KeyDomainStatus, bool)`

GetKeyDomainsOk returns a tuple with the KeyDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDomains

`func (o *HSM) SetKeyDomains(v []KeyDomainStatus)`

SetKeyDomains sets KeyDomains field to given value.


### GetKeys

`func (o *HSM) GetKeys() []Key`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *HSM) GetKeysOk() (*[]Key, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *HSM) SetKeys(v []Key)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


