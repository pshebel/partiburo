import { Suspense } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { Index } from './features/index'
import {Home} from './features/home/home'
import {Guest} from './features/guest/guest'
import {Post} from './features/post/post'
import {Login} from './features/login/login'
import { Unsubscribe } from './features/unsubscribe/unsubscribe'
import { UnsubscribeAll } from './features/unsubscribeAll/unsubscribeAll'
import { Confirm } from './features/confirm/confirm'
import { FullPageLoader } from './components/ui/FullPageLoader'
import { CreateParty } from './features/Party'
import { Announcement } from './features/announcement/announcement'
import { AdminHome } from './features/Admin'
import {ProtectedRoute} from './components/protected-route'

export default function App() {
  return (
    <BrowserRouter>
        <Suspense fallback={<FullPageLoader />}>
          <Routes>
              <Route path="/" element={<Index />}/>
              <Route path="/party" element={<CreateParty />}/>
              <Route path="/:code" element={<Home />} />
              <Route path="/login/:code" element={<Login />} />
              <Route path="/guest/:code" element={<Guest />} />
              <Route path="/post/:code" element={<Post />} />

              {/* admin */}
              <Route path="/admin/:code" element={<AdminHome />} />
              <Route path="/announcement/:code" element={<Announcement />} />
              
              <Route path="/confirm/:code/:passcode" element={<Confirm />} />
              <Route path="/unsubscribe/:party_code/:email_code" element={<Unsubscribe />} />
              <Route path="/unsubscribeAll/:email_code" element={<UnsubscribeAll />} />
          </Routes>
        </Suspense>
    </BrowserRouter>
  )
}
