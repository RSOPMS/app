# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

### Bug Fixes

- #3 update github actions ([f172d04](f172d04315b610ae0ee5f6a642c1df598321d134))
- #2 rename kubernetes app ([9da76f5](9da76f59ccf5584adbaf2969c1131b191023b952))
- #6 fix readme status badges ([a08c95c](a08c95c2a29c65119bc2adde44cd029133546e9f))
- #5 remove sqlite instructions ([4759430](4759430949b8c7c39559b69ca22840f532c7f392))
- #8 changed /health/live to GET /health/live ([030dfdd](030dfddc24e8798eb15e96b769405ae0bda14262))
- #22 update local build context for database ([5ee7311](5ee7311e17fd96531d80a85806c93870d80a4ca3))
- #13 changed list to table ([8c575e5](8c575e53105f4d9d97dba44c4f2112a3d4c55125))
- #13 tidying up code ([e3a8273](e3a8273454f60165a737ab76478e9ed0bd7ea87e))
- #15 comments.tmpl fix ([815298b](815298b050b6d1e0ef5cedc826b166e3a7ca5e00))
- #15 formating ([b67d131](b67d131aaaf6c4b75a4bf5ae6f1bc1588a3337a7))
- #15 formatting ([cd52c31](cd52c31dcad3e51907d2a57b920c793c4331f350))
- #15 formatting ([b8e6ba9](b8e6ba91cbd3b33c477e69b3d009199d3ed9468b))
- #20 start of code cleanup ([9fb982b](9fb982be9276f6110f2602f919c388d2562c1eb1))
- #20 temporary commit ([0613a0a](0613a0a19520cbfefba6969a00714ba40ab93a77))
- #20 final cleanup before review ([19155da](19155dac4ab3567765f765414521382a2210a4d4))
- #10 update static files configuration ([a238bf5](a238bf51119a3deade50722a221c0cc1028145e1))
- #24 add line ending ([a1cb1e9](a1cb1e96e882b0970e0ab40acc18af72f8d29f10))
- #33 update routing ([ba4e767](ba4e76779b705b4435a38fd4a6797b850090c84d))

### Documentation

- #40 add changelog ([9217c0c](9217c0c482f19d512a97ee23eab45a8b4113d74b))

### Features

- #1 add project files ([73e608a](73e608a3800231ef0b92be6e15cf4b9f9554abba))
- #1 add github actions definitions ([7602761](7602761f5394485f0742855d7432c6aae1842695))
- #2 add kubernetes deployment ([651a7fb](651a7fb6e13dcce8005758e0565ad09aa253b723))
- #4 create service for static files ([aa559f0](aa559f08b9ded6dd11b3c9250ddbc30f0e8c0eca))
- #5 swap to postgresql ([9059be7](9059be703e37e471530ca4afc14d1d478f9ae629))
- #6 add instructions and update configurations ([50fc57b](50fc57bca324322e204d504152733933358462ab))
- #12 handle k8s service connections ([4f83b48](4f83b48353c1dd30fff324e76faf636f9b343f35))
- #8 Added health chack handlers for app-static ([6bd6e30](6bd6e309c7f52a2612ab4c3a0437fd4f13b09a09))
- #7 App-issue health check handlers ([bd53e2c](bd53e2cb52d02f570a0aff52634d6cfc18359ba0))
- #17 add webpage layout ([ae75544](ae755440a2dcc676d525b1d6277201b2ed4d4838))
- #21 unify application configuration ([54725e4](54725e4ce9cfdc583dfe9755435841b173ef6f5e))
- #13 project list on /issue/projects, no links to specific project ids yet ([591992b](591992b89dbcf798225be797f7884ae62f7cd55f))
- #13 projects link to /issue/projects/<project_id> ([b934c1a](b934c1a96feba99b2af851f7133767b1e12873e8))
- #14 temp commit ([94887df](94887df255005d85fcd09d89b5486c2f17a608aa))
- #14 links to individual issues ([b6faeed](b6faeeda429d3311e959e228b87356219d9bb9c1))
- #15 Issue page with comments. ([3433c57](3433c578515244f31f1b364ec30e4eb44966e8b1))
- #20 add html headings ([6bcd0cb](6bcd0cb12bc6b26da7f621b9a1326af65b1a7e40))
- #16 k8s health checks ([87f1dab](87f1dabe1d7a155803a06e1310d35e7457dddd14))
- #9 added secrets (postgres db, user, pswd)+makefile ([64b6786](64b6786bc7a7d128e8950a2e5a31e397045cc2d6))
- #10 added configmap, cant debug cuz issues with minikube/docker ([f56b5f2](f56b5f21d5829ba074c64a0f1385e7f7318414d7))
- #10 aligning .env files with what dev branches main.go expects ([f11cce6](f11cce6b16b2e659da102a3f0f2e06c970a82acb))
- #10 import env vars ([90b2247](90b22472ac2895eb9659d197e71d9f52e9c61a53))
- Feat #24 added css styling to table elements ([86ff417](86ff417c1abfff3c3e6e2ac8ebe2008157c52f38))
- Feat #36: Added issue priority in issues table and on issue page ([fb2bffe](fb2bffe181b80265e9471927fa4e3d70668f5b82))
- #39 add editorconfig ([49f5df5](49f5df511d0645893965703c6a1c5042a2ce22a2))

### Refactor

- #20 refactored code ([c0e1013](c0e1013646b894ca3f4c6864b768bb7dee1dba9d))
- #20 add newlines to templates ([1576a9f](1576a9f2c8232937d294a506969d5f140e3704e9))
