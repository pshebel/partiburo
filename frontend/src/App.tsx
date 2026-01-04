import { Suspense } from 'react'
import { Layout } from './features/layout/layout'

export default function App() {
  return (
    <Suspense fallback={<p>Loading...</p>}>
      <Layout />
    </Suspense>
  )
}
