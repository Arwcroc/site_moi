import React from 'react';
// import TwitchPage from './TwitchPage';
import PresentationPage from './PresentationPage';
import Box from '@mui/material/Box';
import { Route, Routes } from "react-router-dom";

const BottomSection = () => {
	return (
		<>
		<Box className="App__WebContainer__BottomSection__PresentationPage"
			sx={{
				width: "100%",
				height: "80%",
			}}
		>
			<Routes>
				<Route path="/" element={
					<>
						<Box className="App__WebContainer__BottomSection__PresentationPage__Instagram">
							<a href= "https://www.instagram.com/arwcroc/?hl=fr" target="_blank" rel="noreferrer">
								<img src="https://upload.wikimedia.org/wikipedia/commons/a/a5/Instagram_icon.png" width={100} height={100}></img>
							</a>
						</Box>
					</>
				}/>
				<Route path="/presentation" element={<PresentationPage/>}/>
			</Routes>
		</Box>
		</>
	);
}

export default BottomSection;