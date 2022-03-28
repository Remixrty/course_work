import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import Navbar from './components/Navbar';
import FirstQ from './pages/FirstQ';
import Home from './pages/Home';
import SecondQ from './pages/SecondQ';
import ThirdQ from './pages/ThirdQ';
// import 

function App() {
  return (
    <>
      <BrowserRouter>
        <Navbar />
        <div className='container pt-4'>
          <Routes>
            <Route path={'/'} exact element={<Home/>} />
            <Route path={'/first'} element={<FirstQ/>} />
            <Route path={'/second'} element={<SecondQ/>} />
            <Route path={'/third'} element={<ThirdQ/>} />

          </Routes>
        </div>
      </BrowserRouter>

    </>
  );
}

export default App;
