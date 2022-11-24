# [1.1.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v1.0.0...v1.1.0) (2022-11-24)


### Bug Fixes

* remove ioutil ([0705b83](https://github.com/IBM/ibm-iae-go-sdk/commit/0705b8301f22a8659f730ab1fb79230a0e317f10))


### Features

* add sdk methods and update existing ([005f927](https://github.com/IBM/ibm-iae-go-sdk/commit/005f92761a7b69ecc06af550b59569c3f33515a7))

# [1.0.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.5.0...v1.0.0) (2022-09-08)


### Features

* add sdk methods for v3 api and update deps ([0134264](https://github.com/IBM/ibm-iae-go-sdk/commit/013426459ba1319e51c1fe53a5f81611a138b04e))


### BREAKING CHANGES

* `CreateInstanceHome` replaced by `SetInstanceHome`

Method `IbmAnalyticsEngineApiV3.CreateInstanceHome` has been replaced by `IbmAnalyticsEngineApiV3.SetInstanceHome`.
`SetInstanceHome` describes the api action better.
Accordingly, options class names have also been updated.

Deprecated: `IbmAnalyticsEngineApiV3.getPlatformLogging` and `IbmAnalyticsEngineApiV3.configurePlatformLogging`

Methods `IbmAnalyticsEngineApiV3.getPlatformLogging` and `IbmAnalyticsEngineApiV3.configurePlatformLogging` are now deprecated.
You can instead use the new methods:
- `IbmAnalyticsEngineApiV3.GetLogForwardingConfig`
- `IbmAnalyticsEngineApiV3.ReplaceLogForwardingConfig`

Signed-off-by: Subin Shekhar <subinpc@gmail.com>

# [0.5.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.5...v0.5.0) (2022-02-11)


### Features

* v3 log config changes ([#17](https://github.com/IBM/ibm-iae-go-sdk/issues/17)) ([3d9e97a](https://github.com/IBM/ibm-iae-go-sdk/commit/3d9e97a3c2d8590a764dd76b5e83d40acdb3c0cd))

## [0.4.5](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.4...v0.4.5) (2022-02-11)


### Bug Fixes

* v3 log config changes   ([2c4b1c7](https://github.com/IBM/ibm-iae-go-sdk/commit/2c4b1c7a03550eab0b1673044b91a87c423c20bc))

## [0.4.4](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.3...v0.4.4) (2021-09-09)


### Bug Fixes

* **analytics engine:** fixed Spark application get api response ([#13](https://github.com/IBM/ibm-iae-go-sdk/issues/13)) ([ec13d54](https://github.com/IBM/ibm-iae-go-sdk/commit/ec13d54b08f2fac72cd0adf899aafc206725a16c))

## [0.4.3](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.2...v0.4.3) (2021-09-09)


### Bug Fixes

* **analytics engine:** fixed request payload for GO Examples ([#12](https://github.com/IBM/ibm-iae-go-sdk/issues/12)) ([c50eddd](https://github.com/IBM/ibm-iae-go-sdk/commit/c50eddd0b412ff2d08dd6c8587a5d4b92f7448a9))

## [0.4.2](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.1...v0.4.2) (2021-09-07)


### Bug Fixes

* **analytics engine:** minor fix to an operation id ([926104f](https://github.com/IBM/ibm-iae-go-sdk/commit/926104fbfc2ab6868e8293aa6295df4eaac2c5e9))

## [0.4.1](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.4.0...v0.4.1) (2021-09-07)


### Bug Fixes

* **analytics engine:** fixed Spark application submit path and payload ([bd231ca](https://github.com/IBM/ibm-iae-go-sdk/commit/bd231ca7ee86e9604ee3e75763a4b9b3fa8eba12))

# [0.4.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.3.1...v0.4.0) (2021-08-11)


### Features

* added support for v3 api's ([59fec0e](https://github.com/IBM/ibm-iae-go-sdk/commit/59fec0ea955106cba3ee5ea8bbb7bf86b3285a24))
* analytics engine v3 feature changes ([4f90700](https://github.com/IBM/ibm-iae-go-sdk/commit/4f90700e52d83d0719ecc8af443bb1de17494936))

# [0.2.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.1.0...v0.2.0) (2021-05-03)


### Features

* **v3 api:** adding v3 api ([7f105a7](https://github.com/IBM/ibm-iae-go-sdk/commit/7f105a744fa0ad351d25c297df1a52c2121c3ffd))

# [0.1.0](https://github.com/IBM/ibm-iae-go-sdk/compare/v0.0.1...v0.1.0) (2020-05-27)


### Features

* **whitelist-api:** Supporting whitelist api ([c4097db](https://github.com/IBM/ibm-iae-go-sdk/commit/c4097db694679b3f416e8f5c17a0f389d0513b8a))
