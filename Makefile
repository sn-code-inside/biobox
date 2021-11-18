progs = al blast2dot cchar cutSeq dnaDist drawKt drawSt fasta2tab getSeq histogram keyMat mum2plot mutator naiveMatcher numAl pam \
plotLine plotSeg plotTree randomizeSeq ranDot ranseq rep2plot repeater revComp rpois sblast shustring simNorm simOrf testMeans var \
watterson wrapSeq
packs = util newick

all:
	test -d bin || mkdir bin
	for prog in $(progs); do \
		make -C $$prog; \
		cp $$prog/$$prog bin; \
	done
	for pack in $(packs); do \
		make -C $$pack; \
	done
.PHONY: doc
doc:
	make -C doc
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
test:
	for prog in $(progs) $(packs); do \
		make test -C $$prog; \
	done

