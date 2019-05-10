import React, { Component } from 'react'

const getPagesArray = (total) => {
    const pagesArray = [];
    let i = 0, remainingTotal = total;
    while (remainingTotal > 0) {
        pagesArray.push(i);
        remainingTotal -= 25;
        i++;
    }
    return pagesArray
}

class Albums extends Component {
    render () {
        const pages = getPagesArray(this.props.total)
        return (
            <section className="albums">
                <ul className="entries">
                {
                    this.props.albums.map(album => {
                        return <li key={ album.id } onClick={ this.props.onClick } data-release-id={ album.id }>{ album.title }</li>
                    })
                }
                </ul>
                <ul className="pages">
                    {
                        pages.map(page => {
                            if (this.props.currentPage === page) {
                                return <li className="current" data-page={ page } key={ page }>{ page+1 }</li>
                            }
                            return <li data-page={ page } key={ page } onClick={ this.props.setPage }>{ page+1 }</li>
                        })
                    }
                </ul>
            </section>
        )
    }
}

export default Albums