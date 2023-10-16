import React from 'react';
import { useState, useEffect } from 'react';
import Avatar from '@mui/material/Avatar';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';

const PresentationPage = () => {
	const [schoolData, setSchoolData] = useState(undefined);
	// const [dataBase, setDataBase] = useState([]);

	useEffect(() => {
		if (schoolData !== undefined) return;
		setSchoolData(undefined);
		fetch("http://localhost:8090/me").then( async (response) => {
			if (response.status < 200 || response.status >= 400) return;
			let body = await response.json();
			setSchoolData(body);
		})
	}, [schoolData]);

	// useEffect(() => {
	// 	if (dataBase !== undefined) return;
	// 	setDataBase(undefined);
	// 	fetch('http://localhost:8080/api/fichier').then( async (response) => {
	// 		if (response.status < 200 || response.status >= 400) return;
	// 		let body = await response.json();
	// 		setDataBase(body);
	// 	})
	// }, [dataBase]);
	
	return (
		<Paper className="App__WebContainer__BottomSection__PresentationPage__CNI"
		elevation={22}
		sx={{
			width:"600px",
			height:"300px",
			backgroundColor:"#cacacaca"
		}}
		>
		{
			schoolData ? (
			<>
			<div className="App__WebContainer__BottomSection__PresentationPage__Avatar">
				<Avatar className="App__WebContainer__BottomSection__PresentationPage__Avatar__Pics"
				sx={{ width: 100, height: 100 }}
				src= {schoolData.User.image.versions.large}
				>
					TF
				</Avatar>
			</div>
			<div className="App__WebContainer__BottomSection__PresentationPage__Info">
				-42 School API-
				<Box className="App__WebContainer__BottomSection__PresentationPage__Name">
					{schoolData.User.first_name} {schoolData.User.last_name}
				</Box>
				<Box className="App__WebContainer__BottomSection__PresentationPage__Login">
					dit {schoolData.User.login}
				</Box>
				<Box className="App__WebContainer__BottomSection__PresentationPage__Level">
					{schoolData.User.Grade} - Level : {schoolData.User.Level}
				</Box>
			</div>
			<div className="App__WebContainer__BottomSection__PresentationPage__Text">
				<Box>
					cacafdfdgdgf sdgd fhdfhg hdfhgd fgfgd fgdfg
				</Box>
			</div>
			</>
			): (
				<></>
			)
		}
		</Paper>
	);
}

export default PresentationPage;