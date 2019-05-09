FOR /L %%A IN (1,1,30) DO (
  set IPFS_PATH=~/.ipfssybil%%A
  START cmd /K sybil-ipfs daemon
)