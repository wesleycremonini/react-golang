import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Root from './routes/root';
import Hello from './routes/hello';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root/>,
    errorElement: <div>Não existeee</div>,
  },
  {
    path: "/hello",
    element: <Hello/>,
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
