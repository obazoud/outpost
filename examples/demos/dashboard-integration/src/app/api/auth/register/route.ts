import { NextRequest, NextResponse } from 'next/server'
import { createUser, getUserByEmail } from '../../../../lib/db'
import { createTenant } from '../../../../lib/outpost'
import { generateTenantId } from '../../../../lib/utils'
import { registerSchema } from '../../../../lib/validations'

export async function POST(request: NextRequest) {
  try {
    const body = await request.json()
    const validatedData = registerSchema.parse(body)

    // Check if user already exists
    const existingUser = await getUserByEmail(validatedData.email)
    if (existingUser) {
      return NextResponse.json(
        { message: 'User already exists' },
        { status: 400 }
      )
    }

    // Create user in database
    const user = await createUser({
      email: validatedData.email,
      name: validatedData.name,
    })

    // Create tenant in Outpost using the user ID
    const tenantId = generateTenantId(user.id.toString())
    console.log(`Attempting to create tenant: ${tenantId} for user: ${user.id}`)
    
    try {
      await createTenant(tenantId)
      console.log(`✅ Tenant created successfully: ${tenantId}`)
      console.log(`🎉 User ${user.id} registered with tenant ${tenantId}`)
    } catch (outpostError) {
      console.error('❌ Failed to create Outpost tenant:', outpostError)
      console.error('❌ Tenant creation error details:', {
        tenantId,
        userId: user.id,
        errorMessage: outpostError?.message,
        errorStack: outpostError?.stack
      })
      // Continue with registration even if tenant creation fails
      // In production, you might want to retry or handle this differently
    }

    return NextResponse.json(
      { message: 'User created successfully' },
      { status: 201 }
    )
  } catch (error: any) {
    console.error('Registration error:', error)
    
    if (error.errors) {
      // Validation error
      return NextResponse.json(
        { message: 'Validation failed', errors: error.errors },
        { status: 400 }
      )
    }

    return NextResponse.json(
      { message: 'Internal server error' },
      { status: 500 }
    )
  }
}