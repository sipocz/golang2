
import pythoncom
import pywintypes
import win32com.client
import win32com.server.util



import openopc2
open_host='localhost'
open_port='7766'
from openopc2.da_client import OpcDaClient
opc = OpcDaClient(open_host, open_port)
opc.connect('Matrikon.OPC.Simulation.1')

aux=opc.list('Simulation Items.Random')
for item in aux:
    print(item)

opc.close()
