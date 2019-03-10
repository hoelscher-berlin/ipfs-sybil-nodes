FOR /L %%A IN (1,1,100) DO (
  set IPFS_PATH=~/.ipfsSybil%%A
  START /B "ipfs%%A" ipfs daemon
)