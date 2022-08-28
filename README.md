# [awesome-apisec](https://github.com/arainho/awesome-apisec)

<h4 align="center">A collection of awesome API Security tools and resources.</h4>
      
<p align="center">
  <a href="#about">About</a> •
  <a href="#api-keys-find-and-validate">API Keys: Find and validate</a> •
  <a href="#books">Books</a> •
  <a href="#cheatsheets">Cheatsheets</a> •
  <a href="#checklist">Checklist</a> •
  <a href="#conferences">Conferences</a> • </br>
  <a href="#deliberately-vulnerable-apis">Deliberately vulnerable APIs</a> •
  <a href="#design-architecture-development">Design, Architecture, Development</a> •
  <a href="#encyclopedias-projects-wikis-and-gitbooks">Encyclopedias, Projects, Wikis and GitBooks</a> • </br>
  <a href="#enumeration-scanning-and-exploration-steps">Enumeration, Scanning and exploration steps</a> •
  <a href="#firewalls">Firewalls</a> •
  <a href="#fuzzing-seclists-wordlists">Fuzzing, SecLists, Wordlists</a> •
  <a href="#http-101">HTTP 101</a> •
  <a href="#mind-maps">Mind maps</a> • </br>
  <a href="#newsletters">Newsletters</a> • 
  <a href="#other-resources">Other resources</a> • 
  <a href="#playlists">Playlists</a> •
  <a href="#podcasts">Podcasts</a> •
  <a href="#presentations-videos">Presentations, Videos</a> •
  <a href="#projects">Projects</a> • </br>
  <a href="#security-apis">Security APIs</a> •
  <a href="#specifications">Specifications</a> •
  <a href="#tools">Tools</a> •
  <a href="#training-workshops-labs">Training, Workshops, Labs</a> •
  <a href="#twitter">Twitter</a> •
  </br>
  • <a href="#contributions">Contributions</a> •
</p>

---

## About
The awesome-api-security (aka awesome-apisec) repository is collection of awesome API Security tools and resources.     
The focus goes to open-source tools and resources that benefit all the community.    

Please read the <a href="#contributions">contributions</a> section before opening a pull request.

## API Keys: Find and validate
| Name | Description |
| ---- | ----------- |
| [API Guesser](https://api-guesser.netlify.app/) | Simple website to guess API Key / OAuth Token by Muhammad Daffa |
|[API Key Leaks: Tools and exploits](https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/API%20Key%20Leaks) | An API key is a unique identifier that is used to authenticate requests associated with your project. Some developers might hardcode them or leave it on public shares. |
| [Key-Checker](https://github.com/daffainfo/Key-Checker)| Go scripts for checking API key / access token validity. |
| [Keyhacks](https://github.com/streaak/keyhacks)| Keyhacks is a repository which shows quick ways in which API keys leaked by a bug bounty program can be checked to see if they're valid. |
|[Private key usage verification ](https://github.com/trufflesecurity/driftwood) | Driftwood is a tool that can enable you to lookup whether a private key is used for things like TLS or as a GitHub SSH key for a user. |

## Books
| Author      | Publisher | Name | Description |
| ----------- | --------- | -----| ----------- |
| Emily Freeman | Data Theorem Special Edition | [API Security for dummies](https://query.prod.cms.rt.microsoft.com/cms/api/am/binary/RWJ9kN) | This book is a high-level introduction to the key concepts of API security and DevSecOps. |
| Neil Madden | Manning   | [API Security in Action](https://www.manning.com/books/api-security-in-action)| API Security in Action teaches you how to create secure APIs for any situation. |
| Dolev Farhi and Nick Aleks | No starch press | [Black Hat GraphQL](https://nostarch.com/black-hat-graphql) | Black Hat GraphQL (book in pre-order). |
| Corey Ball  | No starch press | [Hacking APIs](https://nostarch.com/hacking-apis)| Breaking Web Application Programming Interfaces. |
| Justing Richer and Antonio Sanso | Manning | [Understanding API Security](https://livebook.manning.com/book/understanding-api-security/introduction/) | Several chapters from several Manning books that give you some context for how API security works in the real world. |


## Cheatsheets
| Name | Description |
| ---- | ----------- |
| [GraphQL Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/GraphQL_Cheat_Sheet.html) | GraphQL - OWASP Cheat Sheet Series |
| [JSON Web Token Security Cheat Sheet](https://assets.pentesterlab.com/jwt_security_cheatsheet/jwt_security_cheatsheet.pdf) | PentesterLab - JSON Web Token Security Cheat Sheet |
| [Injection Prevention Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Injection_Prevention_Cheat_Sheet.html) | Injection - OWASP Cheat Sheet Series
| [Microservices Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Microservices_security.html) | Microservices - OWASP Security Cheat Sheet |
| [OWASP API Security Top 10](https://apisecurity.io/encyclopedia/content/owasp-api-security-top-10-cheat-sheet-a4.pdf) | 42Crunch - OWASP API Security Top 10 |
| [REST Assessment Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/REST_Assessment_Cheat_Sheet.html) | REST Assessment - OWASP Cheat Sheet Series |
| [REST Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html) | REST Security - OWASP Cheat Sheet Series |

## Checklist
| Author  | Name | Description |
| ------- | ---- | ------------|
| HolyBugx | [another API Security checklist](https://github.com/HolyBugx/HolyTips/blob/main/Checklist/API%20Security.pdf) | HolyTips: API security checklist
| APIOps Cycles | [API audit checklist](https://www.apiopscycles.com/api-audit-checklist) | API Audit checklist. |
| Shieldfy | [API-Security-Checklist](https://github.com/shieldfy/API-Security-Checklist) | Checklist of the most important security countermeasures when designing, testing, and releasing your API. |
| API Mike, @api_sec | [API penetration testing checklist](https://apimike.com/api-penetration-testing-checklist) | Common steps to include in any API penetration testing process. |
|  Latish Danawale  |  [API Testing Checklist](https://hackanythingfor.blogspot.com/2020/07/api-testing-checklist.html) | API Testing Checklist. |
| Inon Shkedy | [31 days of API Security Tips](https://github.com/smodnix/31-days-of-API-Security-Tips) | This challenge is Inon Shkedy's 31 days API Security Tips. |
| Binary Brotherhood| [OAuth2: Security checklist](https://web.archive.org/web/20210607123429/https://www.binarybrotherhood.io/oauth2_threat_model.html)| OAuth 2.0 Threat Model Pentesting Checklist |
| Apollo | [GraphQL API — GraphQL Security Checklist](https://www.apollographql.com/blog/graphql/security/9-ways-to-secure-your-graphql-api-security-checklist/) | 9 Ways To Secure your GraphQL API — GraphQL Security Checklist |
| LeapGraph | [GraphQL API - The Complete Vulnerability Checklist](https://leapgraph.com/graphql-api-security/)| How to Secure a GraphQL API - The Complete Vulnerability Checklist |
| Lokesh Gupta | [REST API Security Essentials](https://restfulapi.net/security-essentials/) | REST API Tutorial blog entry. |

## Conferences
| Name      | Description |
| --------- | ----------- |
| [APIsecure](https://apisecure.co) | The world's first conference dedicated to API threat management; bringing together breakers, defenders, and solutions in API security. |

## Deliberately vulnerable APIs
| Name | Description |
| ---- | ----------- |
| [APISandbox](https://github.com/API-Security/APISandbox) | Pre-Built Vulnerable Multiple API Scenarios Environments Based on Docker-Compose. |
| [Bookstore](https://tryhackme.com/room/bookstoreoc) | TryHackMe room - A Beginner level box with basic web enumeration and REST API Fuzzing. |
| [crAPI](https://github.com/OWASP/crAPI) | completely ridiculous API (crAPI) |
| [Damn-Vulnerable-GraphQL-Application](https://github.com/dolevf/Damn-Vulnerable-GraphQL-Application)| Damn Vulnerable GraphQL Application is intentionally vulnerable implementation of Facebook's GraphQL technology to learn and practice GraphQL Security. |
| [Damn Vulnerable Micro Services](https://github.com/ne0z/DamnVulnerableMicroServices) | This is a vulnerable microservice written in many languages to demonstrating OWASP API Top Security Risk (under development)|
| [Damn Vulnerable Web Services](https://github.com/snoopysecurity/dvws-node) | Damn Vulnerable Web Services is a vulnerable web service/API/application that we can use to learn webservices/API vulnerabilities. |
| [Generic-University](https://github.com/InsiderPhD/Generic-University) | Vulnerable API with Laravel App |
| [node-api-goat](https://github.com/layro01/node-api-goat) | A simple Express.JS REST API application that exposes endpoints with code that contains vulnerabilities. |
| [Pixi](https://github.com/DevSlop/Pixi) | The Pixi module is a MEAN Stack web app with wildly insecure APIs! |
| [REST API Goat](https://github.com/optiv/rest-api-goat) | This is a "Goat" project so you can get familiar with REST API testing. |
| [VAmPI](https://github.com/erev0s/VAmPI)| Vulnerable REST API with OWASP top 10 vulnerabilities for APIs |
| [vAPI](https://github.com/roottusk/vapi)| vAPI is Vulnerable Adversely Programmed Interface which is Self-Hostable API that mimics OWASP API Top 10 scenarios through Exercises. |
| [vulnapi](https://github.com/tkisason/vulnapi) | Intentionaly very vulnerable API with bonus bad coding practices. |
| [vulnerable-graphql-api](https://github.com/CarveSystems/vulnerable-graphql-api) | A very vulnerable implementation of a GraphQL API. |
| [Websheep](https://github.com/marmicode/websheep) | Websheep is an app based on a willingly vulnerable ReSTful APIs. |

## Design, Architecture, Development
| Name | Description |
| ---- | ----------- |
| [The API Specification Toolbox](http://api.specificationtoolbox.com) | This Toolbox goal is to try and map out all of the different API specifications in use, as well as the services, tooling, extensions, and other supporting elements. |
| [Understanding gRPC, OpenAPI and REST](https://cloud.google.com/blog/products/api-management/understanding-grpc-openapi-and-rest-and-when-to-use-them) | gRPC vs REST: Understanding gRPC, OpenAPI and REST and when to use them in API design |
| [API security design best practices](https://habr.com/en/post/595075/) | API security design best practices for enterprise and public cloud. |
| [REST API Design Guide](https://www.apiopscycles.com/resources/rest-api-design-guide) | This design guide or style guide contains best practices suitable for most REST APIs. |
| [How to design a REST API](https://blog.octo.com/en/design-a-rest-api) | How to design a REST API? - Full guide tackling security, pagination, filtering, versioning, partial answers, CORS, etc.
| [Awesome REST](https://github.com/marmelab/awesome-rest) | A collaborative list of great resources about RESTful API architecture, development, test, and performance. Feel free to contribute to this ongoing list.
| [Collect API Requirements](https://www.apiopscycles.com/collecting-requirements)| Collecting Requirements for your API with APIOps Cycles. |
| [API Audit](https://www.apiopscycles.com/method/api-audit) | API Audit is a method to ensure APIs are matching the API Design guidelines. It also helps check for usability, security and API management platform compatibility. |

## Encyclopedias, Projects, Wikis and GitBooks
| Name | Description |
| ---- | ----------- |
| [APIs Pentest Book](https://pentestbook.six2dez.com/enumeration/webservices/apis) | six2dez - APIs Pentest Book |
| [API Security Empire](https://github.com/cyprosecurity/API-SecurityEmpire) | The API Security Empire Project aims to present unique attack & defense methods in the API Security field |
| [API Security Encyclopedia](https://apisecurity.io/encyclopedia/content/api-security-encyclopedia.htm)  | APIsecurity.io - API Security Encyclopedia |
| [Web API Pentesting](https://book.hacktricks.xyz/pentesting/pentesting-web/web-api-pentesting) | HackTricks - Web API Pentesting |
| [GraphQL](https://book.hacktricks.xyz/network-services-pentesting/pentesting-web/graphql) | HackTricks - GraphQL |

## Enumeration, Scanning and exploration steps
| Name | Description |
| ---- | ----------- |
| [Burp API enumeration](https://portswigger.net/support/using-burp-to-enumerate-a-rest-api) | Using Burp to Enumerate a REST API |
| [ZAP scanning](https://www.zaproxy.org/blog/2017-06-19-scanning-apis-with-zap/) | Scanning APIs with ZAP |
| [ZAP exploring](https://www.zaproxy.org/blog/2017-04-03-exploring-apis-with-zap/)| Exploring APIs with ZAP | 
| [w3af scanning](http://docs.w3af.org/en/latest/scan-rest-apis.html) | Scan REST APIs with w3af |

## Firewalls
| Name | Description |
| ---- | ----------- |
| [Wallarm Free API Firewall](https://github.com/wallarm/api-firewall)| Fast and light-weight API proxy firewall for request and response validation by OpenAPI specs.  |

## Fuzzing, SecLists, Wordlists
| Name | Description |
| ---- | ----------- |
| [API names wordlist](https://github.com/chrislockard/api_wordlist) | A wordlist of API names for web application assessments  |
| [API HTTP requests methods](https://github.com/danielmiessler/SecLists/blob/master/Fuzzing/http-request-methods.txt) | HTTP requests methods  wordlist by @danielmiessler |
| [API Routes Wordlists](https://github.com/assetnote/wordlists/blob/master/data/automated.json) | API Routes - Automated Wordlists provided by Assetnote  |
| [Common API endpoints](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/common-api-endpoints-mazen160.txt) | Wordlist for common API endpoints. |
| [Filenames by fuzz.txt](https://github.com/Bo0oM/fuzz.txt) | Potentially dangerous files | 
| [Fuzzing APIs](https://www.fuzzingbook.org/html/APIFuzzer.html)| Fuzzing APIs chapter from "The Fuzzing Book". 
| [GraphQL SecList](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/graphql.txt) | It's a GraphQL list used during security assessments, collected in one place. |
| [Hacking-APIs](https://github.com/hAPI-hacker/Hacking-APIs) | Wordlists and API paths by @hapi_hacker |
| [Kiterunner Wordlists](https://github.com/assetnote/wordlists/blob/master/data/kiterunner.json) | Kiterunner Wordlists provided by Assetnote |
| [List of API endpoints & objects](https://gist.github.com/yassineaboukir/8e12adefbd505ef704674ad6ad48743d) | A list of 3203 common API endpoints and objects designed for fuzzing. |
| [List of Swagger endpoints](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/swagger.txt) | Swagger endpoints |
| [SecLists for API's web-content discovery](https://github.com/danielmiessler/SecLists/tree/master/Discovery/Web-Content/api) | It is a collection of web content discovery lists for APIs used during security assessments. |

## HTTP 101
| Name | Description |
| ---- | ----------- |
|[Know your HTTP Headers!](http://prezo.s3.amazonaws.com/pixi_california_2018/basics/headers.pdf) | HTTP Headers: a simplified and comprehensive table. |
|[Know your HTTP Methods!](http://prezo.s3.amazonaws.com/pixi_california_2018/basics/methods.pdf) | HTTP Methods: a simplified and comprehensive table. |
|[Know your HTTP Status codes!](http://prezo.s3.amazonaws.com/pixi_california_2018/basics/status-codes.pdf) | HTTP Status codes: a simplified and comprehensive table. |
| [HTTP Status Codes](https://httpstatuses.com/) | httpstatuses.com is an easy to reference database of HTTP Status Codes with their definitions and helpful code references all in one place. |
|[Know your HTTP * Well](https://github.com/for-GET/know-your-http-well)| HTTP headers, media-types, methods, relations and status codes, all summarized and linking to their specification. |

## Mind maps
| Author  | Name | Description |
| ------- | ---- | ----------- |
| [Cypro AB](https://github.com/cyprosecurity) | [API Pentesting - ATTACK](https://github.com/cyprosecurity/API-SecurityEmpire/blob/main/assets/API%20Pentesting%20Mindmap%20ATTACK.pdf) | Mind map: API Pentesting - ATTACK |
| [Cypro AB](https://github.com/cyprosecurity) | [API Pentesting - Recon](https://github.com/cyprosecurity/API-SecurityEmpire/blob/main/assets/API%20Pentesting%20Mindmap.pdf) | Mind map: API Pentesting - Recon |
| [Cypro AB](https://github.com/cyprosecurity) | [GraphQL Attacking](https://github.com/cyprosecurity/API-SecurityEmpire/blob/main/assets/API%20Pentesting%20Mindmap%20%7B%7BGraphQL%20Attacking%7D%7D.pdf) | Mind map: GraphQL Attacking |
| [Mufaddal Masalawala](https://github.com/harsh-bothra) | [IDOR Techniques](https://www.xmind.net/m/CSKSWZ/) | Mind map: IDOR Techniques |
| [David Sopas](https://github.com/dsopas) | [MindAPI](https://dsopas.github.io/MindAPI/play/) | Organize your API security assessment by using MindAPI 
| [Harsh Bothra](https://github.com/muffyhub) | [XML attacks](https://www.xmind.net/m/xNEY9b/) | Mind map: XML attacks |
| [Abhay Bhargav](https://twitter.com/abhaybhargav)| [REST API defenses](https://mobile.twitter.com/abhaybhargav/status/1373982049019654149/photo/1) | Mind map: REST API defenses |

## Newsletters
| Author  | Name | Description |
| ------- | ---- | ----------- |
| 42Crunch | [api security articles](https://apisecurity.io/#newsletter1) | API Security Articles - The Latest API Security News, Vulnerabilities & Best Practices. |

## Other resources
Name | Author | Description |
---- | ------ | ----------- |
| [API Security best practices guide](https://expeditedsecurity.com/api-security-best-practices-megaguide) |  Expedited Security | API Security Best Practices MegaGuide |
| [API Security: The Complete Guide](https://brightsec.com/blog/api-security) | Bright Security | API Security, The Complete Guide |
| [API Penetration Testing](https://blog.securelayer7.net/api-penetration-testing-with-owasp-2017-test-cases) | SecureLayer7 | API Penetration Testing with OWASP 2017 Test Cases. |
|[API Penetration Testing Report](https://underdefense.com/wp-content/uploads/2019/05/Anonymised-API-Penetration-Testing-Report.pdf) | UnderDefense | Anonymised API Penetration Testing Report - vendor sample template |
| [API Pentesting with Swagger Files](https://rhinosecuritylabs.com/application-security/simplifying-api-pentesting-swagger-files/) | RhinoSecurityLabs | Simplifying API Pentesting With Swagger Files. |
| [API security articles](https://char49.com/search/query:API%20security) | Char49 | API security articles. |
| [API Security Testing](https://sphericaldefence.com/api-security-testing) | Spherical Defence | Principles of API Security Testing and how to perform a Security Test on an API. |
| [Finding and Exploiting Web App APIs](https://bendtheory.medium.com/finding-and-exploiting-unintended-functionality-in-main-web-app-apis-6eca3ef000af) | Bend Theory | Finding and Exploiting Unintended Functionality in Main Web App APIs
| [How to Hack an API and Get Away with It](https://smartbear.com/blog/test-and-monitor/api-security-testing-how-to-hack-an-api-part-1/)| SmartBear | How to Hack an API and Get Away with It (Part 1 of 3). |
| [How to Hack APIs in 2021](https://labs.detectify.com/2021/08/10/how-to-hack-apis-in-2021) | Detectify | How to Hack APIs in 2021 |
| [How to Hack API in 60 minutes with Open Source Tools](https://www.wallarm.com/what/how-to-hack-api-in-60-minutes-with-open-source) | Wallarm | How to Hack API in 60 minutes with Open Source Tools | 
| [GraphQL penetration testing](https://blog.yeswehack.com/yeswerhackers/how-exploit-graphql-endpoint-bug-bounty/) | YesWeHAck | How to exploit GraphQL endpoint: introspection, query, mutations & tools. |
| [Fixing the 13 most common GraphQL Vulnerabilities](https://wundergraph.com/blog/the_complete_graphql_security_guide_fixing_the_13_most_common_graphql_vulnerabilities_to_make_your_api_production_ready)| WunderGraph | GraphQL Security Guide, Fixing the 13 most common GraphQL Vulnerabilities to make your API production ready. |
| [Hacking APIs - Notes from Bug Bounty Bootcamp](https://attacker-codeninja.github.io/2021-08-28-Hacking-APIs-notes-from-bug-bounty-bootcamp/)| Aakash Choudhary | My Notes on Hacking APIs from Bug Bounty Bootcamp. |
| [SOAP Security Vulnerabilities and Prevention](https://www.neuralegion.com/blog/top-7-soap-api-vulnerabilities/) | NeuraLegion | SOAP Security, Top Vulnerabilities and How to Prevent Them. |
| [API and microservice security](https://portswigger.net/burp/vulnerability-scanner/api-security-testing/guide-to-api-microservice-security) | PortSwigger | What are API and microservice security? |
| [Strengthening Your API Security Posture](https://42crunch.com/knowledge-series/strengthening-api-security-posture/) | 42Crunch | Strengthening Your API Security Posture – Ford Motor Company. |
| [The Fault in Our Stars](https://www.tenchisecurity.com/blog/thefaultinourstars) | Tenchi Security | Security Implications of AWS API Gateway Lambda Authorizers and IAM Wildcard Expansion. |

## Playlists
| Name | Description |
| ---- | ----------- |
| [Everything API Hacking](https://www.youtube.com/playlist?list=PLbyncTkpno5HqX1h2MnV6Qt4wvTb8Mpol) | A video collection from Katie Paxton-Fear, @InsiderPhD, and other people creating a playlist of API hacking knowledge! |
| [API hacking](https://www.youtube.com/c/TheXSSrat/search?query=API%20hacking)| API hacking videos from @theXSSrat |

## Podcasts
| Name | Description |
| ---- | ----------- |
| [Hacking APIs](https://forallsecure.com/blog/the-hacker-mind-podcast-hacking-apis) | The Hacker Mind Podcast: Hacking APIs |
| [Hack Your API-Security Testing](https://testguild.com/podcast/automation/21-troy-hunt-hack-your-api-security-testing/) | 21: Troy Hunt: Hack Your API-Security Testing. |
| [The OWASP API Security Project](https://podcast.securityjourney.com/erez-yalon-the-owasp-api-security-project/) | Erez Yalon — The OWASP API Security Project |
| [Episode 38 API Security Best Practices](https://wehackpurple.com/podcast/episode-38-api-security-best-practices/) | We Hack Purple Podcast Episode 38 API Security Best Practices. |

## Presentations, Videos
| Name | Description |
| ---- | ----------- |
| [pentesting-rest-apis](https://www.slideshare.net/OWASPdelhi/pentesting-rest-apis-by-gaurang-bhatnagar) | Pentesting Rest API's by Gaurang Bhatnagar |
| [Securing your APIs](https://owasp.org/www-chapter-singapore/assets/presos/Securing_your_APIs_-_OWASP_API_Top_10_2019,_Real-life_Case.pdf) | "How Secure are you APIs?" - Securing your APIs: OWASP API Top 10 2019, Case Study and Demo. |
| [api-security-testing-for-hackers](https://www.bugcrowd.com/resources/webinars/api-security-testing-for-hackers) | API Security Testing For Hackers |
| [bad-api-hapi-hackers](https://www.bugcrowd.com/resources/webinars/bad-api-hapi-hackers)| Bad API, hAPI Hackers! |
| [disclosing-information-via-your-apis](https://www.bugcrowd.com/resources/webinars/hidden-in-plain-site-disclosing-information-via-your-apis/) | Hidden in Plain Site: Disclosing Information via Your APIs. |
| [rest-in-peace-abusing-graphql](https://www.bugcrowd.com/resources/webinars/rest-in-peace-abusing-graphql-to-attack-underlying-infrastructure) | REST in Peace: Abusing GraphQL to Attack Underlying Infrastructure. |

## Projects
| Name | Description |
| ---- | ----------- |
| [owasp api security project](https://owasp.org/www-project-api-security/) | OWASP API Security Project - API Security Top 10 |

## Security APIs
| Name | Description |
| ---- | ----------- |
| [awesome-security-apis](https://github.com/jaegeral/security-apis)| A collective list of public JSON APIs for use in security. |

## Specifications
| Name | Description |
| ---- | ----------- |
| [API Blueprint](https://apiblueprint.org/documentation/specification.html)| API Blueprint Specification | 
| [AscyncAPI](https://www.asyncapi.com/docs/specifications/latest) | AsyncAPI Specification |
| [OpenAPI](https://swagger.io/specification/) | OpenAPI Specification |
| [JSON API](https://jsonapi.org/format/) | JSON API Specification |
| [GraphQL](https://spec.graphql.org/) | GraphQL Specification |
| [RAML](https://github.com/raml-org/raml-spec) | RAML Specification |

## Tools
| Name | Description |
| ---- | ----------- |
| | |
| **GraphQL** |
| [BatchQL](https://github.com/assetnote/batchql) | GraphQL security auditing script with a focus on performing batch GraphQL queries and mutations. |
| [clairvoyance](https://github.com/nikitastupin/clairvoyance) | Obtain GraphQL API schema despite disabled introspection! |
| [InQL](https://github.com/doyensec/inql) | InQL - A Burp Extension for GraphQL Security Testing. |
| [GraphQLmap](https://github.com/swisskyrepo/GraphQLmap)| GraphQLmap is a scripting engine to interact with a graphql endpoint for pentesting purposes. |
| [graphql-path-enum](https://gitlab.com/dee-see/graphql-path-enum) | Tool that lists the different ways of reaching a given type in a GraphQL schema. |
| [graphql-playground](https://github.com/graphql/graphql-playground) | GraphQL IDE for better development workflows (GraphQL Subscriptions, interactive docs & collaboration) |
| [graphql-threat-matrix](https://github.com/nicholasaleks/graphql-threat-matrix) | GraphQL threat framework used by security professionals to research security gaps in GraphQL implementations. |
| [graphw00f](https://github.com/dolevf/graphw00f) | graphw00f is GraphQL Server Engine Fingerprinting utility for software security professionals looking to learn more about what technology is behind a given GraphQL endpoint. |
|  |  |
| **REST APIs** |
| [APICheck](https://bbva.github.io/apicheck/) | The DevSecOps toolset for REST APIs. |
| [APIClarity](https://github.com/apiclarity/apiclarity) | Reconstruct Open API Specifications from real-time workload traffic seamlessly. |
| [APIFuzzer](https://github.com/KissPeter/APIFuzzer) | Fuzz test your application using your OpenAPI or Swagger API definition without coding. |
| [APIKit](https://github.com/API-Security/APIKit) | APIKit：Discovery, Scan and Audit APIs Toolkit All In One. |
| [Arjun](https://github.com/s0md3v/Arjun) | HTTP parameter discovery suite. |
| [Astra](https://github.com/flipkart-incubator/Astra) | Automated Security Testing For REST API's. |
| [Automatic API Attack Tool](https://github.com/imperva/automatic-api-attack-tool) | Imperva's customizable API attack tool takes an API specification as an input, generates and runs attacks that are based on it as an output. |
| [CATS](https://github.com/Endava/cats) | CATS is a REST API Fuzzer and negative testing tool for OpenAPI endpoints. |
| [Cherrybomb](https://github.com/blst-security/cherrybomb) | Stop half-done API specifications with a CLI tool that helps you avoid undefined user behaviour by validating your API specifications. |
| [ffuf](https://github.com/ffuf/ffuf) | Fast web fuzzer written in Go. | 
| [fuzzapi](https://github.com/Fuzzapi/fuzzapi)| Fuzzapi is a tool used for REST API pentesting anTnT-Fuzzerd uses API_Fuzzer gem. |
| [gotestwaf](https://github.com/wallarm/gotestwaf) | An open-source project in Golang to test different web application firewalls (WAF) for detection logic and bypasses |
| [kiterunner](https://github.com/assetnote/kiterunner) | Contextual Content Discovery Tool. |
| [mitmproxy2swagger](https://github.com/alufers/mitmproxy2swagger) | Automagically reverse-engineer REST APIs via capturing traffic |
| [RESTler](https://github.com/microsoft/restler-fuzzer) | RESTler is the first stateful REST API fuzzing tool for automatically testing cloud services through their REST APIs and finding security and reliability bugs in these services. |
| [Swagger-EZ](https://github.com/RhinoSecurityLabs/Swagger-EZ)| A tool geared towards pentesting APIs using OpenAPI definitions. |
| [TnT-Fuzzer](https://github.com/Teebytes/TnT-Fuzzer) | OpenAPI 2.0 (Swagger) fuzzer written in python. Basically TnT for your API. |
| [wadl-dumper](https://github.com/dwisiswant0/wadl-dumper) | Dump all available paths and/or endpoints on WADL file. |
| [fuzz-lightyear](https://github.com/Yelp/fuzz-lightyear)| A pytest-inspired, DAST framework, capable of identifying vulnerabilities in a distributed, micro-service ecosystem through chaos engineering testing and stateful, Swagger fuzzing. |
|  |  |
| **SOAP** |
| [Wsdler](https://github.com/NetSPI/Wsdler)| WSDL Parser extension for Burp. |
| [wsdl-wizard](https://github.com/portswigger/wsdl-wizard)| WSDL Wizard is a Burp Suite plugin written in Python to detect current and discover new WSDL (Web Service Definition Language) files. |
|  |  |
| **Others**|
| [SoapUI](https://github.com/SmartBear/soapui) | SoapUI is a free and open-source cross-platform functional testing solution for APIs and web services. |
| [dredd](https://github.com/apiaryio/dredd)| Language-agnostic HTTP API Testing Tool  |
| [unfurl](https://github.com/tomnomnom/unfurl) | Pull out bits of URLs provided on stdin |

## Training, Workshops, Labs
| Author | Name | Description |
| ------ | ---- | ----------- |
| Pentester Academy | [API security, REST Labs](https://attackdefense.pentesteracademy.com/listing?labtype=rest&subtype=rest-api-security) | Pentester Academy - attack & defense |
| Corey Ball | [API Security University](https://university.apisec.ai) | APIsec University provides training courses for application security professionals |
| Grant Ongers | [API top 10 walkthrough](https://securedelivery.io/articles/api-top-ten-walkthrough/) | OWASP API Top 10 CTF Walk-through. |
| Hacker101 | [GraphQL challenges](https://www.hackerone.com/ethical-hacker/graphql-week-hacker101-capture-flag-challenges) | GraphQL Week on The Hacker101 Capture the Flag Challenges |
| OWASP-SKF | [GraphQL Labs](https://demo.securityknowledgeframework.org/labs/view) | GraphQL Labs on the OWASP Security Knowledge Framework |
| Corey Ball | [Hacking APIs](https://sway.office.com/HVrL2AXUlWGNDHqy) | Hacking APIs: workshop |
| Wesley Thijs | [Let's build an API to hack](https://hackxpert.com/blog/API-Hacking-Excercises/) |  API Hacking Excercises by @TheXSSrat |
| Kontra | [OWASP Top 10 for API](https://application.security/free/owasp-top-10-API) | Is a series of free interactive application security training modules that teach developers how to identify and mitigate security vulnerabilities in their web API endpoints. |
| ShipFast | [Practical API Security Walkthrough](https://github.com/approov/shipfast-api-protection) | Learn practical Mobile and API security techniques: API Key, Static and Dynamic HMAC, Dynamic Certificate Pinning, and Mobile App Attestation. |
| Tushar Kulkarni | [vAPI](https://github.com/roottusk/vapi) | vAPI is Vulnerable Adversely Programmed Interface, Self-Hostable PHP Interface that mimics OWASP API Top 10 scenarios in the means of Exercises. |

## Twitter
| Author  | Name | Description |
| ------- | ---- | ----------- |
| 42Crunch | [@apisecurityio](https://twitter.com/apisecurityio/) | API security news, standards, vulnerabilities, tools. |

## Contributions
1. The purpose of this repository is to collect API Security tools and resources. The preference goes to open-source or community editions tools, creative commons resources, and resources created by the community for the benefit of the community. The exception is in the book's topic, where some referenced items may have an associated cost.

2. Other content vendor-specific, ads, commercial, restricted, free trial, freemium, closed-source (proprietary software), products or services provided in exchange for private user details are considered out of scope pull requests.

3. Duplicated content or entries that do not provide additional or relevant content compared with existing entries may also not be considered.

5. Out of scope pull requests will be probably discarded, closed or ignored without notice.

If you think your content fits the above purposes, please
- create a new branch
- change README.md
- push the new changes
- open a pull request

For more details check GitHub [quickstart/contributing-to-projects](https://docs.github.com/en/get-started/quickstart/contributing-to-projects)
