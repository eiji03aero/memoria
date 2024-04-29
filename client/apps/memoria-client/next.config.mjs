/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['@repo/design-system'],
  output: 'standalone',
};

export default nextConfig;
