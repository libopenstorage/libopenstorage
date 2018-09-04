# Changelog

> NOTE: The SDK is still in tech preview. Once officially released, this changelog will also
> use the SDK version numbers.

## Releases

### v0.6.0 - Tech Preview (8/30/2018)

* Remove unsupported FS Types from supported drivers [#593](https://github.com/libopenstorage/openstorage/issues/593)
* Remove SDK Alert calls as they will be redesinged [#596](https://github.com/libopenstorage/openstorage/issues/596)

### v0.5.0 - Tech Preview (8/25/2018)

* Added `queue_depth` to VolumeSpec and VolumeSpecUpdate
* Remove values from VolumeSpecUpdate which cannot be updated [#590](https://github.com/libopenstorage/openstorage/issues/590)

### v0.4.0 - Tech Preview (8/24/2018)

* Added bucket name and encryption key to SdkCredentialCreateRequest
* Added the ability to disable ssl connection to SdkAwsCredentialRequest

### v0.3.0 - Tech Preview (8/20/2018)

* Added SchedulerNodeName field to StorageNode object

### v0.2.0 - Tech Preview (8/16/2018)

* Changed Credentials.Create to take `name` as a required parameter

### Tech Preview (8/7/2018)

* Added [Identity.Version](https://libopenstorage.github.io/w/generated-api.html#methodopenstorageapiopenstorageidentityversion) Service

### Tech Preview (8/3/2018)

* Added [Idenity](https://libopenstorage.github.io/w/generated-api.html#serviceopenstorageapiopenstorageidentity) Service
* Added [Identity.Capabilities](https://libopenstorage.github.io/w/generated-api.html#methodopenstorageapiopenstorageidentitycapabilities) RPC
