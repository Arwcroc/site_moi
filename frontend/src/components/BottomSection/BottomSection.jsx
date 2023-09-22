import React from 'react';
// import TwitchPage from './TwitchPage';
import PresentationPage from './PresentationPage';
import Box from '@mui/material/Box';


const BottomSection = () => {
	

	return (
		<>
		<Box className="App__WebContainer__BottomSection__PresentationPage"
			sx={{
				width: "100%",
				height: "80%",
			}}
		>
			<PresentationPage/>
		</Box>
			{/* <HobbiesPage/> */}
			{/* <TwitchPage/> */}
			{/* <YoutubePage/> */}
		</>
	);
}

export default BottomSection;