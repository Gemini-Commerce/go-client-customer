version=1.0.0

generate:
	$(MAKE) -C ./modules/go-client-generator/ generate-go-client service=customer version=${version}
push:
	bash git_push.sh
publish:
	GOPROXY=proxy.golang.org go list -m github.com/Gemini-Commerce/go-client-customer@v${version}