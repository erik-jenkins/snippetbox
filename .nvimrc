function GoRunWeb()
  let command = "go run ./cmd/web >infolog.txt 2>errorlog.txt"
  let l:currentWindow=winnr()
  execute "copen"
  exe l:currentWindow . "wincmd w"
  execute "AsyncRun " . command
  sleep 100m
  OpenBrowser http://localhost:4000
endfunction

function GoStopWeb()
  execute "cclose"
  execute "AsyncStop"
endfunction

nnoremap <leader>pr :call GoRunWeb()<CR>
nnoremap <leader>ps :call GoStopWeb()<CR>
