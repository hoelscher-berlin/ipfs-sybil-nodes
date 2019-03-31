FOR /L %%A IN (1,1,20) DO (
  set IPFS_PATH=~/.ipfsSybil%%A
  START cmd /K sybil-ipfs daemon
)