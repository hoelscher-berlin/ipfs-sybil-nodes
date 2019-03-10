FOR /L %%A IN (1,1,30) DO (
  set IPFS_PATH=~/.ipfs%%A
  ipfs init
)

cd %HOMEPATH%
ipfs-sybil-changeconfigs 30