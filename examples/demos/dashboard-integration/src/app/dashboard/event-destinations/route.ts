import { NextRequest, NextResponse } from 'next/server'
import { getToken } from 'next-auth/jwt'
import { getPortalUrl } from '@/lib/outpost'

export async function GET(request: NextRequest) {
  try {
    console.log('=== EVENT DESTINATIONS ROOT API CALLED ===')
    
    // Validate session
    const token = await getToken({ 
      req: request,
      secret: process.env.NEXTAUTH_SECRET 
    })
    
    console.log('Token found:', !!token)
    console.log('Token ID:', token?.id)
    
    if (!token?.id) {
      console.log('No valid token, redirecting to login')
      return NextResponse.redirect(new URL('/auth/login', request.url))
    }

    // Get theme parameter
    const theme = request.nextUrl.searchParams.get('theme') || 'light'
    console.log('Theme:', theme)

    // Generate portal URL using Outpost SDK (user ID = tenant ID)
    const tenantId = token.id as string
    console.log('Generating portal URL for tenant:', tenantId)
    
    const portalUrl = await getPortalUrl(tenantId, theme)
    console.log('Portal URL generated:', portalUrl)

    // Redirect to portal (root path)
    return NextResponse.redirect(portalUrl)
    
  } catch (error) {
    console.error('Portal redirect error:', error)
    return NextResponse.json(
      { error: 'Failed to access event destinations' },
      { status: 500 }
    )
  }
}