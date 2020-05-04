function GoRunWeb()
  let command = "go run ./cmd/web"
  execute "!" . command
endfunction

nnoremap <leader>pr :call GoRunWeb()<CR>
