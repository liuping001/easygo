SUBDIRS=to_redis to_redis_queue

.PHONY: build
build:
	@for subdir in $(SUBDIRS); do \
	    echo "\033[1;33mEntering $$subdir ...\033[m"; \
	    make -C $$subdir; \
	    if test $$? -ne 0; then break; fi; \
	    echo "\033[0;32;32mLeaved $$subdir\033[m"; \
	done

# 清楚生成的程序文件
.PHONY: clean
clean:
	@for subdir in $(SUBDIRS); do \
        echo "\033[1;33mEntering $$subdir ...\033[m"; \
	    make -C $$subdir clean; \
	    if test $$? -ne 0; then break; fi; \
        echo "\033[0;32;32mLeaved $$subdir\033[m"; \
	done