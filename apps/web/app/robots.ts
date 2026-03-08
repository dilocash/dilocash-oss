import { MetadataRoute } from 'next'

export default function robots(): MetadataRoute.Robots {
  return {
    rules: {
      userAgent: '*',
      allow: '/',
      disallow: ['/api/', '/auth/verify/'],
    },
    sitemap: 'https://dilocash.com/sitemap.xml',
  }
}
