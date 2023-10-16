import React from 'react';
import WebContainer from './components/WebContainer.jsx'
import { Route, Routes } from "react-router-dom";
// import PresentationPage from './components/BottomSection/PresentationPage';


function App() {
  return (
	<div className="App">
		<Routes>
			<Route path="/" element={<WebContainer/>}/>
			{/* <Route path="/presentation" element={<PresentationPage/>}/> */}
		</Routes>
    </div>
  );
}

export default App;
