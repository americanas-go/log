# Changelog
All notable changes to this project will be documented in this file.

## [unreleased]

### Bug Fixes

- Changelog commit links ([b2463a4](https://github.com/americanas-go/log/commit/b2463a4f54271ddfb16ef4a32397e7c33773335f))

### Documentation

- Add changelog ([ef0be07](https://github.com/americanas-go/log/commit/ef0be078e1919e70619f0422e576e7dd03a4a97f))

## [v1.6.0](https://github.com/americanas-go/log/compare/v1.5.0...v1.6.0) - 2021-07-12

[16621d0](https://github.com/americanas-go/log/commit/16621d07ed84e3537737e16f7bb1e021b15458ab)...[d873f5b](https://github.com/americanas-go/log/commit/d873f5b91d6a32125bf139449c77be8dbc29c7af)

### Refactor

- Improve package level logging functions ([b97c535](https://github.com/americanas-go/log/commit/b97c535b9e5d641b3b6a2541d756b4f266c4248c))

## [v1.5.0](https://github.com/americanas-go/log/compare/v1.4.0...v1.5.0) - 2021-06-30

[7c7929a](https://github.com/americanas-go/log/commit/7c7929a31a5a960a9d948aa931d33af1494e8cb7)...[16621d0](https://github.com/americanas-go/log/commit/16621d07ed84e3537737e16f7bb1e021b15458ab)

### Documentation

- Add WithError example to README ([51bdea0](https://github.com/americanas-go/log/commit/51bdea04f7cf96f6a7267ac8ff719d11101e42e0))

### Features

- Add WithError method to Logger ([707c08d](https://github.com/americanas-go/log/commit/707c08de9033ac6139e6102aec0def1129f213c2))

### Refactor

- Improve mapToSlice by statically allocating slice size ([fe1dd7f](https://github.com/americanas-go/log/commit/fe1dd7f8ce385d5823616aa17c38ee8b2b0abcdc))

## [v1.4.0](https://github.com/americanas-go/log/compare/v1.3.0...v1.4.0) - 2021-06-11

[3701e44](https://github.com/americanas-go/log/commit/3701e44984b5909d86a05909f6b97de7a26d140c)...[7c7929a](https://github.com/americanas-go/log/commit/7c7929a31a5a960a9d948aa931d33af1494e8cb7)

### Documentation

- Add examples in README.md for logger contract ([d117b97](https://github.com/americanas-go/log/commit/d117b97453fe82a1ffc4ad45c55078059fbbcde4))
- Fix main and logrus README.md ([a0295d6](https://github.com/americanas-go/log/commit/a0295d6ed41af9155cf710595e5eed0b4fd27ecc))
- Adds a pkg.go.dev badge to readme ([e4df3e7](https://github.com/americanas-go/log/commit/e4df3e7c12ed7c454bc3f911e690354429fec10e))
- Fix README.md ([c7074cf](https://github.com/americanas-go/log/commit/c7074cfeb3a1fc4dde17b0c4d2cebaa471345a49))

### Testing

- Fix cloudwatch formatter call ([51de8a5](https://github.com/americanas-go/log/commit/51de8a5dd82cc770c27b8147b8e98ad522755629))

## [v1.3.0](https://github.com/americanas-go/log/compare/v1.2.0...v1.3.0) - 2021-05-24

[713e478](https://github.com/americanas-go/log/commit/713e4785b8846e4129acb61f0533c5eb5b6e8ea8)...[3701e44](https://github.com/americanas-go/log/commit/3701e44984b5909d86a05909f6b97de7a26d140c)

### Bug Fixes

- Rename cloudwatch formatter constructor ([e4a1f1c](https://github.com/americanas-go/log/commit/e4a1f1c45901daae3bbcaa5c0820d2e69aee39be))

### Documentation

- Update README.md ([54778b7](https://github.com/americanas-go/log/commit/54778b7d4c907603de741f12239bbcba0d7b2c80))
- Fix license path ([26afe47](https://github.com/americanas-go/log/commit/26afe47db28f76de246615b1d0cbef82145b1edb))
- Update example in README.md ([47bd2de](https://github.com/americanas-go/log/commit/47bd2def26d9a742f4ca267c304722ce74803816))
- Adds godoc ([95aa40b](https://github.com/americanas-go/log/commit/95aa40b23ee81e7c91ab8da28f3d0f50b71886ca))
- Update comments in logrus option ([bf3c481](https://github.com/americanas-go/log/commit/bf3c4818b2d786aeaeba9f17cedb71b47811c639))
- Add README.md for zap, zerolog and logrus ([83ba3c0](https://github.com/americanas-go/log/commit/83ba3c092c6526f750ad96ccded28b487595b975))
- Update links in README.md ([f94ed42](https://github.com/americanas-go/log/commit/f94ed421a463a1a333b795109991726318e65c16))
- Adds to contribs ([98fe43b](https://github.com/americanas-go/log/commit/98fe43b1965a9c489986b10a442254df0c991065))
- Update zerolog.v1 README.md ([ab1051c](https://github.com/americanas-go/log/commit/ab1051cbbf79b2c172cbc8224be3d19b08a2730c))
- Update zap.v1 README.md ([fd9c31e](https://github.com/americanas-go/log/commit/fd9c31e3733795dc603b760e47ffe8c3eb550cb5))
- Update logrus.v1 README.md ([43c432c](https://github.com/americanas-go/log/commit/43c432c4c3b63723219ed8c72d774ccdb46f7189))
- Add examples to zap, zerolog and logrus ([d6507ed](https://github.com/americanas-go/log/commit/d6507eda91106bdb33c6f837d77510373489769a))
- Add description for logger contract in README.md ([1964192](https://github.com/americanas-go/log/commit/1964192b5cc13248617f3b0c3ce18a37825ae6c7))

### Features

- Add default options for zap ([db110c2](https://github.com/americanas-go/log/commit/db110c2d116465afd5474dbd26a4e4bf0d99bd8f))
- Add default options for zerolog ([958c27d](https://github.com/americanas-go/log/commit/958c27dad9fcf4f117e81f6a68af43f53d2f4d5b))

### Miscellaneous Tasks

- Removes unused func fieldsFromContext ([2a2aa8f](https://github.com/americanas-go/log/commit/2a2aa8f98f87fbc8134d9d397dbcadec16d28205))

<!-- generated by git-cliff -->
