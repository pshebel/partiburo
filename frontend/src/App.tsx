import { Suspense } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import {Home} from './features/home/home'
import {Guest} from './features/guest/guest'
import {Post} from './features/post/post'
import {Login} from './features/login/login'
import { Unsubscribe } from './features/unsubscribe/unsubscribe'
import { UnsubscribeAll } from './features/unsubscribeAll/unsubscribeAll'
import { Confirm } from './features/confirm/confirm'

import {ProtectedRoute} from './components/protected-route'

export default function App() {
  return (
    <BrowserRouter>
        <Suspense fallback={<p>Loading...</p>}>
          <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/confirm/:email/:passcode" element={<Confirm />} />
              <Route path="/unsubscribe/:email" element={<Unsubscribe />} />
              <Route path="/unsubscribeAll/:email" element={<UnsubscribeAll />} />
              <Route 
                path="/" 
                element={
                  <ProtectedRoute>
                    <Home />
                  </ProtectedRoute>
                } 
              />
              <Route 
                path="/guest" 
                element={
                  <ProtectedRoute>
                    <Guest />
                  </ProtectedRoute>
                } 
              />
              <Route 
                path="/post" 
                element={
                  <ProtectedRoute>
                    <Post />
                  </ProtectedRoute>
                } 
              />
          </Routes>
        </Suspense>
    </BrowserRouter>
  )
}
