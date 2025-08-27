import { NextRequest, NextResponse } from 'next/server'
import { getToken } from 'next-auth/jwt'
import { getTenantOverview } from '../../../lib/outpost'

export async function GET(request: NextRequest) {
  try {
    console.log('=== OVERVIEW API CALLED ===')
    
    // Validate session using JWT token
    const token = await getToken({ 
      req: request,
      secret: process.env.NEXTAUTH_SECRET 
    })
    
    if (!token?.id) {
      console.log('No valid token found')
      return NextResponse.json(
        { error: 'Unauthorized' },
        { status: 401 }
      )
    }

    console.log(`Getting overview for user ID: ${token.id}`)
    
    // Use the user ID directly as the tenant ID
    const tenantId = token.id as string
    
    // Get tenant overview from Outpost
    const overview = await getTenantOverview(tenantId)
    
    return NextResponse.json(overview)
  } catch (error) {
    console.error('=== OVERVIEW API ERROR ===')
    console.error('Error:', error)
    
    return NextResponse.json(
      { error: 'Failed to fetch overview data' },
      { status: 500 }
    )
  }
}