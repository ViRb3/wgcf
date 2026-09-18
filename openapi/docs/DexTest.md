# DexTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **string** |  |
**Name** | **string** |  |
**Description** | Pointer to **string** |  | [optional]
**Interval** | **int64** |  |
**Enabled** | **bool** |  |
**Data** | [**DexData**](DexData.md) |  |
**Updated** | **string** |  |
**Created** | **string** |  |

## Methods

### NewDexTest

`func NewDexTest(testId string, name string, interval int64, enabled bool, data DexData, updated string, created string, ) *DexTest`

NewDexTest instantiates a new DexTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexTestWithDefaults

`func NewDexTestWithDefaults() *DexTest`

NewDexTestWithDefaults instantiates a new DexTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *DexTest) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *DexTest) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *DexTest) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetName

`func (o *DexTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DexTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DexTest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DexTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DexTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DexTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DexTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterval

`func (o *DexTest) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DexTest) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DexTest) SetInterval(v int64)`

SetInterval sets Interval field to given value.


### GetEnabled

`func (o *DexTest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DexTest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DexTest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetData

`func (o *DexTest) GetData() DexData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DexTest) GetDataOk() (*DexData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DexTest) SetData(v DexData)`

SetData sets Data field to given value.


### GetUpdated

`func (o *DexTest) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DexTest) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DexTest) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetCreated

`func (o *DexTest) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DexTest) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DexTest) SetCreated(v string)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


