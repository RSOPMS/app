# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

### <!-- 0 -->🚀 Features

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
- #41 add vscode debugging config ([0308f9e](0308f9eb1204a2ff5437f35590a493fe4d43c28e))
- #42 add JWT authentication ([c509821](c509821f04ad84fc30f424d62b5d5f7b33f7087d))
- #29 App-issue healthchecks ping db ([cff18b5](cff18b5fa41d8f9f2bf676824a236aae4dbd5767))
- #35 issue status to issue list (project page) and issue page ([a7d9320](a7d9320f909239920fdaa64421a2d76d683ac312))
- Feat #31: Create new issue form works ([c2e3b83](c2e3b8391b26a44b47eba1d9841b1dace6bd94d2))
- Feat #30: Created form for adding comments ([ba06d48](ba06d48583a9fdd9d1bcd7875944b3cac5965732))
- Feat #32: Added functionality to POST new projects ([3c0390b](3c0390b505d090bfcb701b3522594adcadfe1e82))
- #49 working bulk service skeleton (with healthchecks) ([1576bef](1576bef4878bf439fbf7b82dffa24267bd5d5f7c))
- #49 progress push ([54e7eff](54e7eff729a2b5d2c9d32630b83a00e86f7cd3c5))
- #49 removed unnecesary files ([25628b7](25628b7c768ea8cc7dbc20ea3d107bb02fff257e))
- #49 working with some DB bugs/required improvements ([057444e](057444ea8e97e5c862973df8516be407cf173ed6))
- #49 k8s stuff + makefile (not tested) ([cf8bc8a](cf8bc8ae9db082e3636ec621c1de86c9b71d9a06))
- #49 added project and issue table constraints ([1be0756](1be0756efb2224e4db25ed4433f8ac50afddbe32))
- #49 CI for bulk service ([1a27486](1a274861c5acf90d04b25fc2acbabb8640a7419c))
- #49 README update to include app-bulk status ([c562302](c562302837589153acb6ba92708488a7d8c3de09))
- #49 README update, forgot dev status ([f01d352](f01d352ff395e9f11ecf5b2a7772a71f5c66c6f7))
- #56 added missing data to issue page ([98eeafc](98eeafce623a5c1e57551167bd55ff52f3558164))

### <!-- 1 -->🐛 Bug Fixes

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
- #42 delete go workplace locksum ([daa47c8](daa47c8823774b860f19665ab9b90c8ccd5c841e))
- Fix #31: Refactoring code ([df0979f](df0979f815950916e437e9aa56b92a401b92e918))
- Fix #31: Refactoring code 2 ([8910193](8910193c64d819fcc9036514f2c04d979dfc7d43))
- Fix #30: Fixing URLs and code formatting ([66e8e00](66e8e00f23735cc4d37eb632cf43cb1e71efa9b9))
- #49 bug fixes/validation/constraints ([b4df9dd](b4df9dd3309708abf19a6a92f5d5a5e63dd7b9c2))
- #49 bug fixes for running app-bulk service in k8s (works now) ([3e7ea0b](3e7ea0bd8dd74a2110d0854e2e98956b27fbf51c))
- #49 requisted fixes ([b0fce48](b0fce489cef8789a00dda9e1d8bd1290290663dc))
- #49 fix ([6cb144d](6cb144da4fca62f7e2b1748277b0313fd6402bb4))

### <!-- 10 -->💼 Other

- Merge pull request #11 from RSOPMS/feat/healthcheck ([9f5f810](9f5f81067f8b99bc2a77312bd8b1802ec9df3c3e))
- Merge pull request #18 from RSOPMS/feat/web ([dbe4b6c](dbe4b6c8bbdb8544b5b85660160050022ecb7f57))
- Merge feat/project into prb/feat/project/dev ([46df8c2](46df8c2cbcef15efc6cee9b0ad804ba2bee238fb))
- Merge pull request #23 from RSOPMS/prb/feat/project/dev ([7a2ef21](7a2ef215bcfeb96b30cb8235c8473e35c95265cb))
- Merge feat/k8s into prb/feat/k8s/dev ([e3df412](e3df4120e5108b492c3d34342eb7ec446d65a27d))
- Merge pull request #26 from RSOPMS/prb/feat/k8s/dev ([2e01a6c](2e01a6c7357ed5d177fafb305e65de1768d24ede))
- Merge pull request #27 from RSOPMS/feat/css ([8e6a07c](8e6a07ca256c5e076747c7e455bc71767b41560d))
- Merge pull request #38 from RSOPMS/36-issue-priority ([8ce6948](8ce6948b7dd9220613d3ef22e0dfc07882e3cf16))
- Merge branch '29-healthchecking' into prb/29-healthchecking/dev ([6b5be8d](6b5be8d3c30e4dc288d4073f68770194789f6c8d))
- Merge pull request #46 from RSOPMS/prb/29-healthchecking/dev ([cae6dd3](cae6dd3dca320d597786c19b1d2410b14e1a58b1))
- Merge pull request #44 from RSOPMS/35-issue-status ([80c190f](80c190f65fc7387f6209b1d6dfc9de77aa057446))
- Merge pull request #47 from RSOPMS/37-api-documentation ([78cbdd5](78cbdd571da0aa21e445cb468e3852cbb2ae830c))
- Merge pull request #48 from RSOPMS/feat/add-new-issue ([6de6315](6de6315b1c98e2ab2c90f47fe50d36e7582e81e2))
- Merge pull request #54 from RSOPMS/30-add-new-comment ([eee5043](eee50436a08074d0e79c58d915ddcf9efb7b6ae1))
- Merge pull request #55 from RSOPMS/32-add-new-project ([0140ea7](0140ea7a5bc0638ec58d0140912bf037a5aeb48c))
- Merge pull request #57 from RSOPMS/49-create-bulk-ingress-service ([13fbe23](13fbe236591247ee2c5a3904df96fad1431782b7))
- Merge 56-missing-data into prb/dev/56-missing-data ([7b8ef7c](7b8ef7ccc28efab30bbd5c74ffa1f2ca3feaa4b2))
- Merge pull request #60 from RSOPMS/prb/dev/56-missing-data ([1622959](16229599d0b86121d4dce0ee7aa6e8e218b37055))

### <!-- 2 -->🚜 Refactor

- #20 refactored code ([c0e1013](c0e1013646b894ca3f4c6864b768bb7dee1dba9d))
- #20 add newlines to templates ([1576a9f](1576a9f2c8232937d294a506969d5f140e3704e9))
- #42 update request context ([fd85caf](fd85caf652b9c4230b61e48e803bc06679d11645))

### <!-- 3 -->📚 Documentation

- #40 add changelog ([9217c0c](9217c0c482f19d512a97ee23eab45a8b4113d74b))
- #40 update changelog ([886df82](886df8230dc241d3443f887720a3462932b4ff96))
- #40 add more details to changelog ([770e577](770e5770c574b591d83d5f3e6c04876b6f31c20a))
- #41 add debugging documentation ([b92c157](b92c157b5dc109742aa81cac4bdd2f7170f35d78))
- #41 fix typo ([218aa2b](218aa2b432469156a1f67624284c43cbbc42281a))
- #37 OpenAPI doc .yaml ([6a1fd2f](6a1fd2fa412886870fde36f8129d615b18fa3c25))
- #37 fixed paths in OpenAPI docs .yaml ([db9755c](db9755c9e2c794a16ca2a35c40c82b665d8462f1))
- #37 fixed responses in OpenAPI docs .yaml ([99e938d](99e938d1590f2dfa41d9d9b9ceba3014ead043cf))
- #37 openapi formatting ([22b0757](22b07573e9db4c95a73b7619d949a6a53b9c032a))
- #37 fixed responses in OpenAPI docs .yaml ([197a5a5](197a5a5b1aa36e9417e1495a5a1e8daa23d5dd17))

### <!-- 5 -->🎨 Styling

- #42 prepare package for multiple keys ([b7aa39c](b7aa39c0bb9c00446f593d47a17fbaf2549ef730))
- #35 unify indentation style ([e5fe187](e5fe187a169be4c4c5b0817b391a07f696ea801c))
- #31 update modal code ([46022bd](46022bd6eb5d65f4e2e952792fb5b3f39e9f592f))
- #32 code formatting and update changelog ([e578ae7](e578ae7f99c0215949495e588ecdf8098e509493))
- #56 code formatting ([d3a17a6](d3a17a62eaf452061499ff178a439efe2f592da4))


