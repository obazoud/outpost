import { NextRequest, NextResponse } from 'next/server'
import { getToken } from 'next-auth/jwt'
import { getPortalUrl } from '@/lib/outpost'

export async function GET(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  try {
    console.log('=== EVENT DESTINATIONS CATCH-ALL API CALLED ===')
    
    // Validate session
    const token = await getToken({ 
      req: request,
      secret: process.env.NEXTAUTH_SECRET 
    })
    
    console.log('Token found:', !!token)
    console.log('Token ID:', token?.id)
    console.log('Path segments:', params.path)
    
    if (!token?.id) {
      console.log('No valid token, redirecting to login')
      return NextResponse.redirect(new URL('/auth/login', request.url))
    }

    // Extract and validate the path
    const pathSegments = params.path || []
    const portalPath = pathSegments.length > 0 ? '/' + pathSegments.join('/') : ''
    
    console.log('Portal path:', portalPath)
    
    // Validate the path - allow /new and /destinations/{uuid}
    const isNewRoute = portalPath === '/new'
    const isDestinationRoute = !!portalPath.match(/^\/destinations\/[a-f0-9\-]{36}$/)
    const isValidPath = isNewRoute || isDestinationRoute
    
    if (!isValidPath) {
      console.log('Invalid path, redirecting to root portal')
      return NextResponse.redirect(new URL('/dashboard/event-destinations', request.url))
    }

    // Get theme parameter
    const theme = request.nextUrl.searchParams.get('theme') || 'light'
    console.log('Theme:', theme)

    // Generate portal URL using Outpost SDK (user ID = tenant ID)
    const tenantId = token.id as string
    console.log('Generating portal URL for tenant:', tenantId)
    
    const portalUrl = await getPortalUrl(tenantId, theme)
    console.log('Portal URL generated:', portalUrl)

    // Append the portal path to the portal URL
    const finalUrl = new URL(portalUrl)
    console.log('Original portal pathname:', finalUrl.pathname)
    
    if (portalPath) {
      // Handle path construction carefully to avoid double slashes
      const basePath = finalUrl.pathname.endsWith('/') ? finalUrl.pathname.slice(0, -1) : finalUrl.pathname
      finalUrl.pathname = basePath + portalPath
      console.log('Constructed pathname:', finalUrl.pathname)
    }
    
    console.log('Final URL with path:', finalUrl.toString())

    // Redirect to portal
    return NextResponse.redirect(finalUrl.toString())
    
  } catch (error) {
    console.error('Portal redirect error:', error)
    return NextResponse.json(
      { error: 'Failed to access event destinations' },
      { status: 500 }
    )
  }
}