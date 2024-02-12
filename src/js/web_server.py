#!/usr/bin/env python3
import http.server
from http.server import SimpleHTTPRequestHandler, HTTPServer
import socketserver

class CORSHTTPRequestHandler(SimpleHTTPRequestHandler):
    def end_headers(self):
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header('Access-Control-Allow-Methods', 'GET, POST, OPTIONS')
        self.send_header('Cache-Control', 'no-store, no-cache, must-revalidate')
        return super(CORSHTTPRequestHandler, self).end_headers()

if __name__ == '__main__':
    port = 8000
    with HTTPServer(('', port), CORSHTTPRequestHandler) as httpd:
        print(f"Serving HTTP on 0.0.0.0 port {port} (http://0.0.0.0:{port}/) ...")
        httpd.serve_forever()
