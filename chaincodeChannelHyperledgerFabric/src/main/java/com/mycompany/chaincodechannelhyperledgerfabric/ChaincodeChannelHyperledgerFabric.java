package com.mycompany.chaincodechannelhyperledgerfabric;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.contract.Contract;
import org.hyperledger.fabric.shim.ChaincodeStub;

@Contract
public class ChaincodeChannelHyperledgerFabric implements ContractInterface {

    @Override
    public Context createContext(ChaincodeStub stub) {
        return new Context(stub);
    }

    public void criarAtivo(Context ctx, String ativoID, String proprietario) {
        // Obtenha o registro do ativo do estado do contrato
        Ativo ativo = ctx.getAtivoRegistro().get(ativoID);

        // Verifique se o ativo já existe
        if (ativo != null) {
            throw new RuntimeException("O ativo com ID " + ativoID + " já existe.");
        }

        // Crie um novo ativo
        Ativo novoAtivo = new Ativo(ativoID, proprietario);

        // Coloque o novo ativo no estado do contrato
        ctx.getAtivoRegistro().put(ativoID, novoAtivo);
    }

    public String consultarAtivo(Context ctx, String ativoID) {
        // Obtenha o registro do ativo do estado do contrato
        Ativo ativo = ctx.getAtivoRegistro().get(ativoID);

        // Verifique se o ativo existe
        if (ativo == null) {
            throw new RuntimeException("O ativo com ID " + ativoID + " não existe.");
        }

        // Retorne os detalhes do ativo
        return "Detalhes do ativo com ID: " + ativoID + "\nProprietário: " + ativo.getProprietario();
    }

    public void atualizarAtivo(Context ctx, String ativoID, String novoProprietario) {
        // Obtenha o registro do ativo do estado do contrato
        Ativo ativo = ctx.getAtivoRegistro().get(ativoID);

        // Verifique se o ativo existe
        if (ativo == null) {
            throw new RuntimeException("O ativo com ID " + ativoID + " não existe.");
        }

        // Atualize o proprietário do ativo
        ativo.setProprietario(novoProprietario);

        // Atualize o ativo no estado do contrato
        ctx.getAtivoRegistro().put(ativoID, ativo);
    }
    
    public static void main(String[] args) {
        System.out.println("Hello World!");
    }
}
