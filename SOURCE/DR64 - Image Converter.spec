# -*- mode: python ; coding: utf-8 -*-


a = Analysis(
    ['C:\\Users\\danre\\OneDrive\\Desktop\\DR64 Software\\OPEN SOURCE\\DR64 - Image Converter\\DR64 - Image Converter.py'],
    pathex=[],
    binaries=[],
    datas=[],
    hiddenimports=[],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    noarchive=False,
    optimize=0,
)
pyz = PYZ(a.pure)

exe = EXE(
    pyz,
    a.scripts,
    a.binaries,
    a.datas,
    [],
    name='DR64 - Image Converter',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    upx_exclude=[],
    runtime_tmpdir=None,
    console=False,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
    icon=['C:\\Users\\danre\\OneDrive\\Desktop\\DR64 Software\\OPEN SOURCE\\DR64 - Image Converter\\DR64 - Image Converter.ico'],
)
