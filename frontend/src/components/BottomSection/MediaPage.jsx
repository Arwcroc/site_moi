import React from 'react';
import { useState, useEffect } from 'react';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Typography from '@mui/material/Typography';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';

const TwitchElement = (props) => {
	const [twitchDataJDG, setTwitchDataJDG] = useState();

	const fetchTwitchJDG = () => {
		if (twitchDataJDG !== undefined) return;
		console.log("fetching api twitch " + props.username)
		fetch('http://localhost:8090/twitchapi?user_id=' + props.username).then( async (response) => {
			if (response.status < 200 || response.status >= 400) {
				console.error("Error on fetch api twitch " + props.username)
			};
			let body = await response.json();
			// console.log(body)
			console.log("got api twitch " + props.username)
			setTwitchDataJDG(body);
		})
	}
	useEffect(() => {
		fetchTwitchJDG()
		// eslint-disable-next-line
	}, [setTwitchDataJDG]);

	return (
		<Grid item xs={10} md={5}>
			<Box className="App__WebContainer__BottomSection__MediaPage__TwitchPlace__jdg">
				<Grid container spacing={2}>
					<Grid item>		
						<Avatar src={twitchDataJDG && twitchDataJDG.user.profile_image_url}>JDG</Avatar> 
					</Grid>
					<Grid item xs={12} sm container>
						<Grid item xs container direction="column" spacing={2}>
							<Grid item xs>
								<Typography gutterBottom variant="subtitle1" component="div">
									{twitchDataJDG && twitchDataJDG.user.display_name}
								</Typography>
								<Button>No LIVE</Button>
							</Grid>
						</Grid>
					</Grid>
				</Grid>
			</Box>				
		</Grid>
	)
}

const MediaPage = () => {


	return (
		<div className="App__WebContainer__BottomSection__MediaPage">
			<div className="App__WebContainer__BottomSection__MediaPage__TwitchPlace">
				<Grid 
				container
				columnSpacing={8}
				rowSpacing={6}
				direction="row"
				justifyContent="center"
				alignItems="center"
				>
					<TwitchElement username="joueur_du_grenier" />
					<TwitchElement username="zerator" />
					<TwitchElement username="antoinedaniel" />
					<TwitchElement username="mistermv" />
				</Grid>
			</div>
			<div className="App__WebContainer__BottomSection__MediaPage__YoutubePlace">
				
			</div>
		</div>
	);
}

export default MediaPage;