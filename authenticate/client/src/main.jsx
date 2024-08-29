import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import SignInPage from './pages/SignInPage.jsx';
import ProfilePage from './pages/ProfilePage.jsx';
const router = createBrowserRouter([
  {
    path: "/",
    element: <ProfilePage/>,
  },
  {
    path: "/sign-in",
    element: <SignInPage/>,
  },
]);
ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
      <RouterProvider router={router} />
    {/* <App /> */}
  </React.StrictMode>,
)
