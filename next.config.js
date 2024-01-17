/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'export',
  async rewrites() {
    return {
      fallback: [
        {
          source: '/api/:path*',
          destination: `http://localhost:9999/api/:path*`,
        },
        {
          source: '/socket.io',
          destination: `http://localhost:9999/socket.io/`,
        }
      ],
    }
  },
  reactStrictMode: true,
}

module.exports = nextConfig
