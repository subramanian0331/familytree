Family Tree
===
Abstract: This is a Family Tree Application written in Golang. It uses RedisGraph to save the genealogy data.


## Install & Dependence
- golang
- redisgraph
- postgres

## Dataset Preparation
| Dataset | Download |
| ---     | ---   |
| dataset-A | [download]() |
| dataset-B | [download]() |
| dataset-C | [download]() |

## Use
- for train
  ```
  python train.py
  ```
- for test
  ```
  python test.py
  ```
## Pretrained model
| Model | Download |
| ---     | ---   |
| Model-1 | [download]() |
| Model-2 | [download]() |
| Model-3 | [download]() |


## Directory Hierarchy
```
|—— api
|    |—— server.go
|—— cache
|—— go.mod
|—— go.sum
|—— handlers
|    |—— authHandlers.go
|    |—— handlers.go
|—— main.go
|—— models
|    |—— member.go
|    |—— relationship_jsonenums.go
|    |—— sex_jsonenums.go
|    |—— user.go
|—— store
|    |—— postgresDB.go
|    |—— redisGraphDB.go
|    |—— store.go
|—— ui
|    |—— templates
|        |—— index.html
|        |—— success.html

```
## Code Details
### Tested Platform
- software
  ```
  OS: Debian unstable (May 2021), Ubuntu LTS
  Python: 3.8.5 (anaconda)
  PyTorch: 1.7.1, 1.8.1
  ```
- hardware
  ```
  CPU: Intel Xeon 6226R
  GPU: Nvidia RTX3090 (24GB)
  ```
### Hyper parameters
```
```
## References
- [paper-1]()
- [paper-2]()
- [code-1](https://github.com)
- [code-2](https://github.com)
  
## License

## Citing
If you use xxx,please use the following BibTeX entry.
```
```
