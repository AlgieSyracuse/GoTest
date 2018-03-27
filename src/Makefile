# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
TARGET = hello

SOURCE = hello.go

BINDIR = ../Debug

GOFLAGS =

.PHONY : all build clean run

all : clean build

build :
	mkdir -p $(BINDIR)
	$(GOBUILD) -o $(BINDIR)/$(TARGET) -v

clean : 
	rm -rf $(BINDIR)/*

run :
	$(BINDIR)/$(TARGET)