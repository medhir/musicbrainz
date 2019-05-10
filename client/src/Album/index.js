import React, { Component, Fragment } from 'react'
import CoverArt from './CoverArt'

const AlbumView = (props) => {
    const trackList = props.album.Release.MediumList.Mediums[0].TrackList.Tracks;
    return (
        <Fragment>
            <CoverArt id={ props.id }/>
            <TrackList tracks={ trackList }/>
        </Fragment>
    )
}

const TrackList = (props) => {
    return (
        <ul>
        {
            props.tracks.map(track => <li>{ track.Recording.Title }</li>)
        }
        </ul>
    )
}

class Album extends Component {
    render () {
        return (
            <section className="album">
                <AlbumView album={ this.props.album } id={ this.props.album.Release.id }/>
            </section>
        )
    }
}

export default Album;