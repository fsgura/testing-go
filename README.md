<h1>TESTING-GO</h1>

<i>A repository just to start a set of tests, using a base default go installation on linux</i>

<b>Testing OS</b>: <i>linux</i><br/>
<b>Distribution</b>: <i>Fedora 25</i><br/>
<b>Workspace</b>: <i>Go 1.7</i><br/>
installed using
</br>
<code>dnf -y install go-golang</code>
</br>
and dependencies from the standard fedora repo

<h2>NOTE</h2>
Fedora 25 provides a dnf set of packages for a ready to use go workspace.
By the way, when installing the bundle the environment remains enigmatic, especially when considering the correct setup of the Environment.<br/>
<br/>
An initial look at the packages filesystem structure can give and idea of how the go provided toolset has been structured.<br/>
<br/>
<b>Main path:</b><br/>
<code>/usr/lib/golang</code>
<i>In this folder there are all the subfolders with the main components of the development tool</i>

