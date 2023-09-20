import React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';

const MenuSection = () => {
	return (
		<Box className="App__WebContainer__MenuSection"
			sx={{
				width: "100%",
				height: "20%",
				background: "#cacadada",
				borderBottom: "solid 1px darkgrey",
			}}
		>
			<Box className="App__WebContainer__MenuSection__Home">
				<Typography variant="h4">Téo Froissart</Typography>
			</Box>
			<Box className="App__WebContainer__MenuSection__Home__Page">
				<Box className="App__WebContainer__MenuSection__Home__Page__Presentation">
					<Typography variant="button">Présentation</Typography>
				</Box>
				<Box className="App__WebContainer__MenuSection__Home__Page__Hobbies">
					<Typography variant="button">Hobbies</Typography>
				</Box>
				<Box className="App__WebContainer__MenuSection__Home__Page__Twitch">
					<Typography variant="button">Twitch</Typography>
				</Box>
				<Box className="App__WebContainer__MenuSection__Home__Page__Youtube">
					<Typography variant="button">Youtube</Typography>
				</Box>
			</Box>
		</Box>
	);
}

export default MenuSection;