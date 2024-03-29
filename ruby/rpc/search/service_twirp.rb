# Code generated by protoc-gen-twirp_ruby 1.3.0, DO NOT EDIT.
require 'twirp'
require_relative 'service_pb.rb'

module Github
  module Minimaltwirp
    module Rpc
      class SearchService < Twirp::Service
        package 'github.minimaltwirp.rpc'
        service 'Search'
        rpc :Search, SearchRequest, SearchResponse, :ruby_method => :search
      end

      class SearchClient < Twirp::Client
        client_for SearchService
      end
    end
  end
end
