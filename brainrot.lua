vim.lsp.start({
	name = "brainrot-lsp",
	cmd = { "./bin/brainrot-lsp" },
	root_dir = vim.fn.getcwd(),
})
