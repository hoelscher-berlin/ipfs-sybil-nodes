FOR /L %%A IN (1,1,1) DO (
  set IPFS_PATH=~/.ipfsddos%%A
  START cmd /K ddos2-ipfs daemon
)