FOR /L %%A IN (1,1,40) DO (
  set IPFS_PATH=~/.ipfsddos%%A
  START cmd /K ddos2-ipfs log level dht error
)