import { Pool } from 'pg'

const pool = new Pool({
  connectionString: process.env.DATABASE_URL,
  max: 20,
  idleTimeoutMillis: 30000,
  connectionTimeoutMillis: 2000,
})

export interface User {
  id: string
  name?: string | null
  email: string
  emailVerified?: Date | null
  image?: string | null
  createdAt: Date
  updatedAt: Date
}

export async function getUserTenant(userId: string): Promise<User | null> {
  try {
    const result = await pool.query(
      'SELECT id, name, email, "emailVerified", image, "createdAt", "updatedAt" FROM users WHERE id = $1',
      [userId]
    )
    return result.rows[0] || null
  } catch (error) {
    console.error('Error fetching user tenant:', error)
    return null
  }
}

export async function createUser(userData: {
  email: string
  name?: string
  hashedPassword?: string
}): Promise<User> {
  try {
    const result = await pool.query(
      `INSERT INTO users (email, name, "createdAt", "updatedAt") 
       VALUES ($1, $2, NOW(), NOW()) 
       RETURNING id, name, email, "emailVerified", image, "createdAt", "updatedAt"`,
      [userData.email, userData.name || null]
    )
    return result.rows[0]
  } catch (error) {
    console.error('Error creating user:', error)
    throw error
  }
}

export async function getUserByEmail(email: string): Promise<User | null> {
  try {
    const result = await pool.query(
      'SELECT id, name, email, "emailVerified", image, "createdAt", "updatedAt" FROM users WHERE email = $1',
      [email]
    )
    return result.rows[0] || null
  } catch (error) {
    console.error('Error fetching user by email:', error)
    return null
  }
}

export default pool