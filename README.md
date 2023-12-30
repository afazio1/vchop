<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">vchop</h3>

  <p align="center">
    A quick and dirty CLI tool that removes silences from videos.
    <br />
    <br />
    ·
    <a href="https://github.com/afazio1/vchop/issues">Report Bug</a>
    ·
    <a href="https://github.com/afazio1/vchop/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Removing pauses and silence is a mundane task for any video editor. `vchop` automates this task through an elegant and user-friendly CLI tool. Powered by FFmpeg, Golang, and the terminal UI framework Bubble Tea, `vchop` transforms your video files, leaving them free of unnecessary pauses.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
### Built With

* [![Golang][Golang]][Golang-url]
* [![FFmpeg][FFmpeg]][FFmpeg-url]
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To install `vchop` locally, follow these steps.

### Prerequisites

`vchop` relies on FFmpeg for silence detection and removal. You can install it from the official page [here](https://www.ffmpeg.org/download.html)
### Installation

1. Install the binary
   ```sh
   go install github.com/afazio1/vchop
   ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage
### Flags
- `-input` the video file you want to remove silence from
- `-output` the edited video file you want to output
- `-noise` the maximum noise threshold (dB) that is to be considered silence
- `-duration` the minimum duration threshold (seconds) that is to be considered silence
### Interactive UI
```sh
vchop create
```
- Enter flag values as prompted by the UI
### Command Only
```sh
vchop -input=video.mov -output=output.mp4 -noise=-30 -duration=0.5
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* Interactive UI and code structure inspired by [Melkey's Go Blueprint](https://github.com/Melkeydev/go-blueprint)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/afazio1/vchop.svg?style=for-the-badge
[contributors-url]: https://github.com/afazio1/vchop/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/afazio1/vchop.svg?style=for-the-badge
[forks-url]: https://github.com/afazio1/vchop/network/members
[stars-shield]: https://img.shields.io/github/stars/afazio1/vchop.svg?style=for-the-badge
[stars-url]: https://github.com/afazio1/vchop/stargazers
[issues-shield]: https://img.shields.io/github/issues/afazio1/vchop.svg?style=for-the-badge
[issues-url]: https://github.com/afazio1/vchop/issues
[license-shield]: https://img.shields.io/github/license/afazio1/vchop.svg?style=for-the-badge
[license-url]: https://github.com/afazio1/vchop/blob/master/LICENSE.txt
[product-screenshot]: images/screenshot.png
[Golang]: https://img.shields.io/badge/Golang-29BEB0?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev
[FFmpeg]: https://img.shields.io/badge/FFmpeg-gray?style=for-the-badge&logo=ffmpeg&logoColor=green
[FFmpeg-url]: https://ffmpeg.org
