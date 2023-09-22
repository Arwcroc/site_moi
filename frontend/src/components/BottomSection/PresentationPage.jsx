import React from 'react';
import { useState, useEffect } from 'react';
import Avatar from '@mui/material/Avatar';
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
		<Box className="App__WebContainer__BottomSection__PresentationPage__CNI"
		elevation={15}
		sx={{
			width:"600px",
			height:"300px",
		}}
		>
		{
			schoolData ? (
			<div className="App__WebContainer__BottomSection__PresentationPage__Avatar">
				<Avatar className="App__WebContainer__BottomSection__PresentationPage__Avatar__Pics"
				sx={{ width: 100, height: 100 }}
				src= {schoolData.User.image.versions.large}
				>
					H
				</Avatar>
			</div>
			): (
				<></>
			)
		}
		</Box>
	);
}

export default PresentationPage;