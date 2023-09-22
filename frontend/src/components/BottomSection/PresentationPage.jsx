import React from 'react';
import { useState, useEffect } from 'react';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
// import Typography from '@mui/material/Typography';


const PresentationPage = () => {
	const [schoolData, setSchoolData] = useState(undefined);

	useEffect(() => {
		if (schoolData !== undefined) return;
		setSchoolData(undefined);
		fetch("http://localhost:8090/me").then( async (response) => {
			if (response.status < 200 || response.status >= 400) return;
			let body = await response.json();
			setSchoolData(body);
		})
	}, [schoolData]);

	return (
		<Paper className="App__WebContainer__BottomSection__PresentationPage_CNI"
		elevation={15}
		sx={{
			width:"600px",
			height:"300px",
		}}
		>
			"coucou"
			<Box className="App__WebContainer__BottomSection__PresentationPage_Pics">
				Coucou
			</Box>
			{/* schoolData ?
			<Typography variant="h5">{schoolData.User.Login}</Typography> */}
		</Paper>
	);
}

export default PresentationPage;