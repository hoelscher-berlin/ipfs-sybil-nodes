FOR /L %%A IN (1,1,30) DO (
  set IPFS_PATH=~/.ipfs%%A
  START /B "ipfs%%A" ipfs daemon
)